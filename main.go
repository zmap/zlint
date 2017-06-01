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
	inPath         string
	outPath        string
	multi          bool
	threaded       bool
	numCertThreads int
	prettyPrint    bool
	numProcs       int
	channelSize    int
)

func init() {
	flag.StringVar(&inPath, "in-file", "", "File path for the input certificate(s).")
	flag.StringVar(&outPath, "out-file", "-", "File path for the output JSON.")
	flag.BoolVar(&prettyPrint, "list-lints-json", false, "Use this flag to print supported lints in JSON format, one per line")
	flag.BoolVar(&multi, "multi", false, "Use this flag to specify inserting many certs at once. Certs in this mode must be Base64 encoded DER strings, one per line.")
	flag.BoolVar(&threaded, "threads", false, "Use this flag to specify that -multi mode runs multi-threaded. This has no effect otherwise.")
	flag.IntVar(&numCertThreads, "cert-threads", 1, "Use this flag to specify the number of threads in -threads mode.  This has no effect otherwise.")
	flag.IntVar(&numProcs, "procs", 1, "Use this flag to specify the number of processes to run on.")
	flag.IntVar(&channelSize, "channel-size", 10000, "Use this flag to specify the number of values in the buffered channel.")
	flag.Parse()
}

func ProcessCertificate(in <-chan string, wg *sync.WaitGroup) {
	log.Info("Processing certificates...")
	defer wg.Done()
	for raw := range in {
		der, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			//Handle error
		}
		parsed, err := x509.ParseCertificate(der)
		if err != nil { //could not parse
			//log.Info("Could not parse certificate with error:", err)
		} else { //parsed
			zlintResult := zlint.ZLintResultTestHandler(parsed)
			jsonResult, err := json.Marshal(zlintResult.ZLint)
			if err != nil {
				log.Fatal("Could not parse JSON.")
			}
			fmt.Println(string(jsonResult))
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

func main() {
	log.SetLevel(log.InfoLevel)
	runtime.GOMAXPROCS(numProcs)

	if prettyPrint {
		zlint.PrettyPrintZLint()
		return
	}

	//Initialize Channel
	certs := make(chan string, channelSize)
	var readerWG sync.WaitGroup
	var procWG sync.WaitGroup
	readerWG.Add(1)
	go ReadCertificate(certs, inPath, &readerWG)
	for i := 0; i < numCertThreads; i++ {
		procWG.Add(1)
		go ProcessCertificate(certs, &procWG)
	}
	go func() {
		readerWG.Wait()
		close(certs)
	}()
	procWG.Wait()
}
