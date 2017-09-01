/*
 * ZLint Copyright 2017 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint"
)

var ( //flags
	listLintsJSON bool
	prettyprint   bool
	format        string
)

func init() {
	flag.BoolVar(&listLintsJSON, "list-lints-json", false, "Use this flag to print supported lints in JSON format, one per line")
	flag.StringVar(&format, "format", "pem", "One of {pem, der, base64}")
	flag.BoolVar(&prettyprint, "pretty", false, "Pretty-print output")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] file...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	log.SetLevel(log.InfoLevel)
}

func main() {

	if listLintsJSON {
		zlint.EncodeLintDescriptionsToJSON(os.Stdout)
		return
	}

	var inform = strings.ToLower(format)
	if flag.NArg() < 1 || flag.Arg(0) == "-" {
		lint(os.Stdin, inform)
	} else {
		for _, filePath := range flag.Args() {
			var inputFile *os.File
			var err error
			inputFile, err = os.Open(filePath)
			if err != nil {
				log.Fatalf("unable to open file %s: %s", filePath, err)
			}
			var fmt = inform
			switch {
			case strings.HasSuffix(filePath, ".der"):
				fmt = "der"
			case strings.HasSuffix(filePath, ".pem"):
				fmt = "pem"
			}
			lint(inputFile, fmt)
			inputFile.Close()
		}
	}
}

func lint(inputFile *os.File, inform string) {
	fileBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatalf("unable to read file %s: %s", inputFile.Name(), err)
	}

	var asn1Data []byte
	switch inform {
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
	jsonBytes, err := json.Marshal(zlintResult.Results)
	if err != nil {
		log.Fatalf("unable to encode lints JSON: %s", err)
	}
	if prettyprint {
		var out bytes.Buffer
		if err := json.Indent(&out, jsonBytes, "", " "); err != nil {
			log.Fatalf("can't format output: %s", err)
		}
		os.Stdout.Write(out.Bytes())
	} else {
		os.Stdout.Write(jsonBytes)
	}
	os.Stdout.Write([]byte{'\n'})
	os.Stdout.Sync()
}
