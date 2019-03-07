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
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint"
	"github.com/zmap/zlint/lints"
)

var ( // flags
	listLintsJSON   bool
	listLintsSchema bool
	prettyprint     bool
	format          string
	include         string
	exclude         string
)

func init() {
	flag.BoolVar(&listLintsJSON, "list-lints-json", false, "Print supported lints in JSON format, one per line")
	flag.BoolVar(&listLintsSchema, "list-lints-schema", false, "Print supported lints as a ZSchema")
	flag.StringVar(&format, "format", "pem", "One of {pem, der, base64}")
	flag.StringVar(&include, "include", "", "Comma-separated list of lints to include by name")
	flag.StringVar(&exclude, "exclude", "", "Comma-separated list of lints to exclude by name")

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

	if listLintsSchema {
		names := make([]string, 0, len(lints.Lints))
		for lintName := range lints.Lints {
			names = append(names, lintName)
		}
		sort.Strings(names)
		fmt.Printf("Lints = SubRecord({\n")
		for _, lintName := range names {
			fmt.Printf("    \"%s\":LintBool(),\n", lintName)
		}
		fmt.Printf("})\n")
		return
	}

	// include/exclude lints based on flags
	setLints()

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
			var fileInform = inform
			switch {
			case strings.HasSuffix(filePath, ".der"):
				fileInform = "der"
			case strings.HasSuffix(filePath, ".pem"):
				fileInform = "pem"
			}

			lint(inputFile, fileInform)
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

func setLints() {
	if include != "" && exclude != "" {
		log.Fatal("unable to use include and exclude flag at the same time")
	}

	includeLints()
	excludeLints()
}

func includeLints() {
	if include == "" {
		return
	}

	// parse includes to map for easier matching
	var includes = strings.Split(include, ",")
	var includesMap = make(map[string]struct{}, len(includes))
	for _, includeName := range includes {
		includeName = strings.TrimSpace(includeName)
		if _, ok := lints.Lints[includeName]; !ok {
			log.Fatalf("unknown lint %q in include list", includeName)
		}

		includesMap[includeName] = struct{}{}
	}

	// clear all initialised lints except for includes
	for lintName := range lints.Lints {
		if _, ok := includesMap[lintName]; !ok {
			delete(lints.Lints, lintName)
		}
	}
}

func excludeLints() {
	if exclude == "" {
		return
	}

	// parse excludes to map to get rid of duplicates
	var excludes = strings.Split(exclude, ",")
	var excludesMap = make(map[string]struct{}, len(excludes))
	for _, excludeName := range excludes {
		excludesMap[strings.TrimSpace(excludeName)] = struct{}{}
	}

	// exclude lints
	for excludeName := range excludesMap {
		if _, ok := lints.Lints[excludeName]; !ok {
			log.Fatalf("unknown lint %q in exclude list", excludeName)
		}

		delete(lints.Lints, excludeName)
	}
}
