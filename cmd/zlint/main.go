package main

import (
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint"
)

var ( //flags
	listLintsJSON bool
	format        string
)

func init() {
	flag.BoolVar(&listLintsJSON, "list-lints-json", false, "Use this flag to print supported lints in JSON format, one per line")
	flag.StringVar(&format, "format", "pem", "One of {pem, der, base64}")
	flag.Parse()

	log.SetLevel(log.InfoLevel)
}

func main() {

	if listLintsJSON {
		zlint.EncodeLintDescriptionsToJSON(os.Stdout)
		return
	}

	var inputFile *os.File
	if flag.NArg() < 1 || flag.Arg(0) == "-" {
		inputFile = os.Stdin
	} else {
		filePath := flag.Arg(0)
		var err error
		inputFile, err = os.Open(filePath)
		if err != nil {
			log.Fatalf("unable to open file %s: %s", filePath, err)
		}
	}

	fileBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatalf("unable to read file %s: %s", inputFile.Name(), err)
	}

	var asn1Data []byte
	switch inform := strings.ToLower(format); inform {
	case "pem":
		p, _ := pem.Decode(fileBytes)
		if p == nil || p.Type != "CERTIFICATE" {
			log.Fatal("unable to parse PEM")
		}
		asn1Data = p.Bytes
	case "der":
		asn1Data = fileBytes
	case "base64":
		asn1Data, err = base64.StdEncoding.DecodeString(string(fileBytes))
		if err != nil {
			log.Fatalf("unable to parse base64: %s", err)
		}
	default:
		log.Fatalf("unknown input format %s", format)
	}

	c, err := x509.ParseCertificate(asn1Data)
	if err != nil {
		log.Fatalf("unable to parse certificate: %s", err)
	}

	zlintResult := zlint.LintCertificate(c)
	jsonBytes, err := json.Marshal(zlintResult.ZLint)
	if err != nil {
		log.Fatalf("unable to encode lints JSON: %s", err)
	}
	os.Stdout.Write(jsonBytes)
	os.Stdout.Write([]byte{'\n'})
	os.Stdout.Sync()
}
