/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/formattedoutput"
	"github.com/zmap/zlint/v3/lint"

	_ "github.com/zmap/zlint/v3/profiles"
)

var ( // flags
	listLintsJSON   bool
	listLintSources bool
	listProfiles    bool
	summary         bool
	longSummary     bool
	prettyprint     bool
	format          string
	nameFilter      string
	includeNames    string
	excludeNames    string
	includeSources  string
	excludeSources  string
	profile         string
	printVersion    bool
	config          string
	exampleConfig   bool

	// version is replaced by GoReleaser or `make` using an LDFlags option at
	// build time. Here we supply a default value for folks that `go install` or
	// `go build` directly from src.
	version = "dev-unknown"
)

func init() {
	flag.BoolVar(&listLintsJSON, "list-lints-json", false, "Print lints in JSON format, one per line")
	flag.BoolVar(&listLintSources, "list-lints-source", false, "Print list of lint sources, one per line")
	flag.BoolVar(&listProfiles, "list-profiles", false, "Print profiles in JSON format, one per line")
	flag.BoolVar(&summary, "summary", false, "Prints a short human-readable summary report")
	flag.BoolVar(&longSummary, "longSummary", false, "Prints a human-readable summary report with details")
	flag.StringVar(&format, "format", "pem", "One of {pem, der, base64}")
	flag.StringVar(&nameFilter, "nameFilter", "", "Only run lints with a name matching the provided regex. (Can not be used with -includeNames/-excludeNames)")
	flag.StringVar(&includeNames, "includeNames", "", "Comma-separated list of lints to include by name")
	flag.StringVar(&excludeNames, "excludeNames", "", "Comma-separated list of lints to exclude by name")
	flag.StringVar(&includeSources, "includeSources", "", "Comma-separated list of lint sources to include")
	flag.StringVar(&excludeSources, "excludeSources", "", "Comma-separated list of lint sources to exclude")
	flag.StringVar(&profile, "profile", "", "Name of the linting profile to use. Equivalent to enumerating all of the lints in a given profile using includeNames")
	flag.BoolVar(&printVersion, "version", false, "Print ZLint version and exit")
	flag.StringVar(&config, "config", "", "A path to valid a TOML file that is to service as the configuration for a single run of ZLint")
	flag.BoolVar(&exampleConfig, "exampleConfig", false, "Print a complete example of a configuration that is usable via the '-config' flag and exit. All values listed in this example will be set to their default.")

	flag.BoolVar(&prettyprint, "pretty", false, "Pretty-print JSON output")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ZLint version %s\n\n", version)
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] file...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	log.SetLevel(log.InfoLevel)
}

//nolint:cyclop
func main() {
	if printVersion {
		fmt.Printf("ZLint version %s\n", version)
		return
	}

	// Build a registry of lints using the include/exclude lint name and source
	// flags.
	registry, err := setLints()
	if err != nil {
		log.Fatalf("unable to configure included/exclude lints: %v\n", err)
	}

	if listLintsJSON {
		registry.WriteJSON(os.Stdout)
		return
	}

	if exampleConfig {
		b, err := registry.DefaultConfiguration()
		if err != nil {
			log.Fatalf("a critical error occurred while generating a configuration file, %s", err)
		}
		fmt.Println(string(b))
		return
	}

	if listLintSources {
		sources := registry.Sources()
		sort.Sort(sources)
		for _, source := range sources {
			fmt.Printf("    %s\n", source)
		}
		return
	}

	if listProfiles {
		enc := json.NewEncoder(os.Stdout)
		enc.SetEscapeHTML(false)
		for _, profile := range lint.AllProfiles() {
			err = enc.Encode(profile)
			if err != nil {
				log.Fatalf("a critical error occurred while JSON encoding a profile, %s", err)
			}
		}
		return
	}

	var inform = strings.ToLower(format)
	if flag.NArg() < 1 || flag.Arg(0) == "-" {
		doLint(os.Stdin, inform, registry)
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

			doLint(inputFile, fileInform, registry)
			inputFile.Close()
		}
	}
}

//nolint:cyclop
func doLint(inputFile *os.File, inform string, registry lint.Registry) {
	fileBytes, err := io.ReadAll(inputFile)
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

	zlintResult := zlint.LintCertificateEx(c, registry)
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
		fmt.Printf("\n\n")
	}
	if summary {
		formattedoutput.OutputSummary(zlintResult, false)
	}
	if longSummary {
		formattedoutput.OutputSummary(zlintResult, true)
	}
	if !prettyprint && !summary && !longSummary {
		os.Stdout.Write(jsonBytes)
	}
	os.Stdout.Write([]byte{'\n'})
	os.Stdout.Sync()
}

// trimmedList takes a comma separated string argument in raw, splits it by
// comma, and returns a list of the separated elements after trimming spaces
// from each element.
func trimmedList(raw string) []string {
	var list []string
	for _, item := range strings.Split(raw, ",") {
		list = append(list, strings.TrimSpace(item))
	}
	return list
}

// setLints returns a filtered registry to use based on the nameFilter,
// includeNames, excludeNames, includeSources, and excludeSources flag values in
// use.
//
//nolint:cyclop
func setLints() (lint.Registry, error) {
	configuration, err := lint.NewConfigFromFile(config)
	if err != nil {
		return nil, err
	}
	lint.GlobalRegistry().SetConfiguration(configuration)
	// If there's no filter options set, use the global registry as-is
	anyFilters := func(args ...string) bool {
		for _, arg := range args {
			if arg != "" {
				return true
			}
		}
		return false
	}
	if !anyFilters(nameFilter, includeNames, excludeNames, includeSources, excludeSources, profile) {
		return lint.GlobalRegistry(), nil
	}
	filterOpts := lint.FilterOptions{}
	if nameFilter != "" {
		r, err := regexp.Compile(nameFilter)
		if err != nil {
			return nil, fmt.Errorf("bad -nameFilter: %v", err)
		}
		filterOpts.NameFilter = r
	}
	if excludeSources != "" {
		if err := filterOpts.ExcludeSources.FromString(excludeSources); err != nil {
			log.Fatalf("invalid -excludeSources: %v", err)
		}
	}
	if includeSources != "" {
		if err := filterOpts.IncludeSources.FromString(includeSources); err != nil {
			log.Fatalf("invalid -includeSources: %v\n", err)
		}
	}
	if excludeNames != "" {
		filterOpts.ExcludeNames = trimmedList(excludeNames)
	}
	if includeNames != "" {
		filterOpts.IncludeNames = trimmedList(includeNames)
	}
	if profile != "" {
		p, ok := lint.GetProfile(profile)
		if !ok {
			return nil, fmt.Errorf("lint profile name does not exist: %v", profile)
		}
		filterOpts.AddProfile(p)
	}

	return lint.GlobalRegistry().Filter(filterOpts)
}
