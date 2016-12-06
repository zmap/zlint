//certlint contains functions to check certificates for standards complience.
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/zmap/zlint/lints"
	"github.com/zmap/zlint/zlint"
	"github.com/zmap/zlint/zlint/ringbuff"
)

const CHUNKSIZE int = 10000 //number of certs per work unit, must be >=1
const THREADS int = 4       //number of processing threads for --threads mode, must be >=1

var ( //flags
	inPath   string
	outPath  string
	outStat  string
	multi    bool
	threaded bool
)

var ( //sync values for --threads
	inBuffer      ringbuff.RingBuffer
	outBuffer     ringbuff.RingBuffer
	poisonBarrier sync.WaitGroup //used prevent outBuffer from being poisoned before Enqueueing is complete
	//barrier to prevent the early main thread exits
	exittex       sync.Mutex //guards writeComplete and mainWait
	writeComplete bool       //output complete if true
	mainWait      *sync.Cond //condition vairable for the main thread, associated with exittex
)

func init() {
	flag.StringVar(&inPath, "in-file", "", "File path for the input certificate(s).")
	flag.StringVar(&outPath, "out-file", "-", "File path for the output JSON.")
	flag.StringVar(&outStat, "out-stat", "-", "File path for the output stats.")
	flag.BoolVar(&multi, "multi", false, "Use this flag to specify inserting many certs at once. Certs in this mode must be Base64 encoded DER strings, one per line.")
	flag.BoolVar(&threaded, "threads", false, "Use this flag to specify that --multi mode runs multi-threaded. This has no effect otherwise.")
	flag.Parse()
}

func main() {
	var err error
	var theJSON []byte
	if multi { //multiple cert file, do chunk reading method or multi-threaded method
		if threaded {
			threadMode()
		} else {
			err = multiMode() //write happens internally
			if err != nil {
				fmt.Println(err)
			}
		}
		return //no further work
	} else { //single cert, read whole file
		theJSON, err = singleMode()
	}

	if err != nil {
		fmt.Println("Something went wrong....")
		fmt.Println(err)
	}

	if outPath != "" && outPath != "-" {
		ioutil.WriteFile(outPath, theJSON, 0644)
	} else {
		fmt.Println(string(theJSON))
	}

}

func multiMode() (err error) {
	//reader variables
	var fileReader io.Reader //base file reader, to be buffered
	fileReader, err = os.Open(inPath)
	if err != nil { //file open/read error
		return err
	}
	var buffReader *bufio.Reader = bufio.NewReader(fileReader) //buffered file reader from base reader
	//writer
	var fileWriter *os.File
	if outPath != "" && outPath != "-" { //only open file if path specified
		fileWriter, err = os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666) //os.File in write mode
		if err != nil {                                                                  //file open/write error
			return err
		}
		defer fileWriter.Close()
	}
	m := make(map[string]int)
	//main computation
	var done bool = false
	for !done { //loops until read error, consider reworking this
		certs := make([]string, 0, CHUNKSIZE)                          //input buffer
		reports := make([]map[string]lints.ResultStruct, 0, CHUNKSIZE) //output buffer
		var readIn, lintOut int = 0, 0                                 //reading and processing counters
		for ; readIn < CHUNKSIZE && !done; readIn++ {                  //read in CHUNKSIZE # certs
			lineIn, err := buffReader.ReadString('\n')
			if err != nil { //read error, stop reading and process anything that was read
				done = true
				if err != io.EOF {
					fmt.Println(err)
					break //if not eof, assume incomplete cert read
				}
			}
			if len(lineIn) > 0 { //dont add empty line at the end
				certs = append(certs, lineIn) //add cert to process queue
			}
		}

		//process read certs

		for x := 0; x < len(certs); x++ {
			reportOut, err := zlint.Lint64(certs[x], m) //lint the cert
			if err != nil {
				fmt.Println(err)
			} else {
				reports = append(reports, reportOut) //move report to out buffer
				lintOut++
			}
		}

		//output results
		var fileWriter2 *os.File
		if outStat != "" && outStat != "-" {
			fileWriter2, err = os.OpenFile(outStat, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
			for i, j := range m {
				fileWriter2.WriteString(i + " : " + strconv.Itoa(j) + "\n")
			}
			fileWriter2.Close()
		}

		for x := 0; x < len(reports); x++ {
			theJSON, err := json.Marshal(reports[x])
			if err != nil {
				fmt.Println(err)
				continue //didn't marshal, continue
			}
			if outPath != "" && outPath != "-" {
				_, err := fileWriter.Write(theJSON)
				if err != nil { //write failure
					fmt.Println(err)
					return err
				}
				_, err = fileWriter.WriteString("\n")
				if err != nil { //write failure
					fmt.Println(err)
					return err
				}
			} else { //output to stdout
				fmt.Println(string(theJSON))
			}
		}

	}

	return nil //if it reaches this point, no fatal errors have occurred
}

func threadMode() {
	//initialize thread sync variables
	inBuffer.Init(30)
	outBuffer.Init(40)
	mainWait = sync.NewCond(&exittex)
	poisonBarrier.Add(1 + THREADS) //requires all processing threads AND the reader to Done()

	//initiate reader
	go readChunks()

	//initiate processing
	for i := 1; i <= THREADS; i++ {
		go processChunks()
	}

	//initiate writer
	go writeChunks()

	//entering the poison barrier
	poisonBarrier.Wait() //wait for all processing threads and the reader to finish
	//exiting poison barrier

	outBuffer.Poison() //notify output thread of processing complete

	//entering the exit barrier
	exittex.Lock()

	for !writeComplete {
		mainWait.Wait() //requires exittex
	}

	exittex.Unlock()
	//exiting the exit barrier

	//program execution complete, only the main thread still running
	return
}

// reader thread function for threaded --multi mode
func readChunks() {
	//set-up reader variables
	var fileReader io.Reader //base file reader, to be buffered
	fileReader, err := os.Open(inPath)
	if err != nil { //file open/read error
		fmt.Println(err)
		os.Exit(1) //Fatal Error: Abort!
	}
	var buffReader *bufio.Reader = bufio.NewReader(fileReader) //buffered file reader from base reader

	//begin reading
	var done bool = false
	for !done {
		certs := make([]string, 0, CHUNKSIZE) //local input buffer

		for readIn := 0; readIn < CHUNKSIZE && !done; readIn++ { //read in CHUNKSIZE # certs
			lineIn, err := buffReader.ReadString('\n')
			if err != nil { //read error, stop reading
				done = true
				if err != io.EOF {
					fmt.Println(err)
					break //if not eof, assume incomplete cert read
				}
			}
			if len(lineIn) > 0 { //dont add empty line at the end
				certs = append(certs, lineIn) //add cert to local input buffer
			}
		}

		//chunk read in
		if len(certs) != 0 { //don't add empty jobs to the queue
			inBuffer.Enqueue(certs) //add chunk to process queue; this function can block
		}

	}
	//reading complete/aborted, release hold on main thread and poison the process queue
	inBuffer.Poison()
	poisonBarrier.Done()
}

// cert processing worker thread function for threaded --multi mode
func processChunks() {
	//var lintOut int = 0//local processed counter (unused)
	m := make(map[string]int)
	for true {
		var chunk []string = inBuffer.Dequeue()             //get job from work queue
		var reports []string = make([]string, 0, CHUNKSIZE) //output buffer
		if chunk == nil {
			//queue poisoned, prepare to exit thread
			break
		}

		for x := 0; x < len(chunk); x++ { //process read certs
			reportOut, err := zlint.Lint64(chunk[x], m) //lint the cert, reportOut is a map[string]zlint.ResultStruct
			if err != nil {
				fmt.Println(err)
			} else {
				//convert cert to string for output
				stringOut, err := json.Marshal(reportOut)
				if err != nil {
					fmt.Println(err)
				} else {
					reports = append(reports, string(stringOut)) //move report to out buffer
					//lintOut++ (unused)
				}
			}
		}
		//write to outbuffer
		outBuffer.Enqueue(reports)
	}
	//thread exiting, release hold on main thread
	poisonBarrier.Done()
}

// writer thread function for threaded --multi mode
func writeChunks() {
	//set-up writer variables
	var fileWriter *os.File
	var err error
	if outPath != "" && outPath != "-" { //only open file if path specified
		fileWriter, err = os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666) //os.File in write mode
		if err != nil {                                                                  //file open/write error
			fmt.Println(err)
			os.Exit(2) //Fatal Error: Abort
		}
	}

	for true { //loop until poisoned queue return
		outStrings := outBuffer.Dequeue() //[]string
		if outStrings == nil {
			//queue poioned and empty
			break //stop writing and prepare for thread exit
		}

		//print/write all strings in output job
		for _, x := range outStrings {
			if outPath != "" && outPath != "-" { //output to file set
				_, err := fileWriter.WriteString(x)
				if err != nil { //write failure
					fmt.Println(err)
					os.Exit(2) //Fatal Error: Abort
				}
				_, err = fileWriter.WriteString("\n")
				if err != nil { //write failure
					fmt.Println(err)
					os.Exit(2) //Fatal Error: Abort
				}
			} else { //output to stdout set
				fmt.Println(x)
			}
		}
	}
	//writeing complete, clean-up and release main thread to exit
	fileWriter.Close()

	//entering main thread exit barrier
	exittex.Lock()

	writeComplete = true
	mainWait.Signal() //wake main thread in case it beat us here

	//exiting main thread exit barrier
	exittex.Unlock()
}

func singleMode() ([]byte, error) {

	theCert := lints.ReadCertificate(inPath)

	if theCert == nil {
		return nil, errors.New("Parsing Failed")
	}
	m := make(map[string]int)
	theReport, err := zlint.ParsedTestHandler(theCert, m)
	if err != nil {
		return nil, err
	}

	return json.Marshal(theReport)
}
