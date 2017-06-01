package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/zlint"
	"os"
	"runtime"
	"sync"
)

var ( //flags
	inPath           string
	outPath          string
	numCertThreads   int
	prettyPrint      bool
	numProcs         int
	channelSize      int
	crashIfParseFail bool
)

func init() {
	flag.StringVar(&inPath, "input-file", "", "File path for the input certificate(s).")
	flag.StringVar(&outPath, "output-file", "-", "File path for the output JSON.")
	flag.BoolVar(&prettyPrint, "list-lints-json", false, "Use this flag to print supported lints in JSON format, one per line")
	flag.IntVar(&numCertThreads, "cert-threads", 1, "Use this flag to specify the number of threads in -threads mode.  This has no effect otherwise.")
	flag.IntVar(&numProcs, "procs", 0, "Use this flag to specify the number of processes to run on.")
	flag.IntVar(&channelSize, "channel-size", 10000, "Use this flag to specify the number of values in the buffered channel.")
	flag.BoolVar(&crashIfParseFail, "crash", false, "Use this flag if you want to crash on parsing failure.")
	flag.Parse()
}

func ProcessCertificate(in <-chan string, out chan<- []byte, wg *sync.WaitGroup) {
	log.Info("Processing certificates...")
	defer wg.Done()
	for raw := range in {
		der, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			//Handle error
		}
		parsed, err := x509.ParseCertificate(der)
		if err != nil { //could not parse
			if crashIfParseFail {
				log.Fatal("Could not parse certificate with error: ", err)
			} else {
				log.Info("Could not parse certificate with error: ", err)
			}
		} else { //parsed
			zlintResult := zlint.ZLintResultTestHandler(parsed)
			jsonResult, err := json.Marshal(zlintResult.ZLint)
			if err != nil {
				log.Fatal("Could not parse JSON.")
			}
			out <- jsonResult
		}
	}
}

func ReadCertificate(out chan<- string, filename string, wg *sync.WaitGroup) {
	log.Info("Reading certificates...")
	defer wg.Done()
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if err = scanner.Err(); err != nil {
			log.Fatal("Error with scanning file: ", err)
		}
	} else {
		log.Fatal("Error reading file: ", err)
	}
}

func WriteOutput(in <-chan []byte, outputFileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	var outFile *os.File
	var err error
	var writeOutputToFile bool = false
	if outputFileName != "" && outputFileName != "-" {
		outFile, err = os.Create(outputFileName)
		if err != nil {
			log.Fatal("Unable to create output file: ", err)
		}
		defer outFile.Close()
		writeOutputToFile = true
	}
	for json := range in {
		if writeOutputToFile {
			outFile.Write(json)
			outFile.Write([]byte{'\n'})
		} else {
			fmt.Println(string(json))
		}
	}
}

func main() {
	log.SetLevel(log.InfoLevel)
	runtime.GOMAXPROCS(numProcs)

	if prettyPrint {
		zlint.PrettyPrintZLint()
		return
	}

	//Initialize Channels
	certs := make(chan string, channelSize)
	jsonOut := make(chan []byte, channelSize)

	var readerWG sync.WaitGroup
	var procWG sync.WaitGroup
	var writerWG sync.WaitGroup

	readerWG.Add(1)
	writerWG.Add(1)

	go ReadCertificate(certs, inPath, &readerWG)
	go WriteOutput(jsonOut, outPath, &writerWG)

	for i := 0; i < numCertThreads; i++ {
		procWG.Add(1)
		go ProcessCertificate(certs, jsonOut, &procWG)
	}

	go func() {
		readerWG.Wait()
		close(certs)
	}()

	procWG.Wait()
	close(jsonOut)
	writerWG.Wait()
}
