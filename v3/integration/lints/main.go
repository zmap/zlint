package main

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

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zmap/zlint/v3/integration/lints/lint"
	"github.com/zmap/zlint/v3/integration/lints/lints"
)

var linters = []lint.Lint{
	&lints.InitFirst{},
	&lints.NotCommittingGenTestCerts{},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("USAGE %s <path to lint directory>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	results, err := run(os.Args[1])
	if err != nil {
		fmt.Printf("A fatal error has occurred: %v\n", err)
		os.Exit(2)
	}
	if len(results) == 0 {
		os.Exit(0)
	}
	fmt.Printf("Found %d linting errors\n", len(results))
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
	os.Exit(1)
}

func run(dir string) ([]*lint.Result, error) {
	var results []*lint.Result
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !isAGoFile(info) {
			return nil
		}
		r, err := lint.RunLints(path, linters)
		if err != nil {
			return err
		}
		results = append(results, r...)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func isAGoFile(info os.FileInfo) bool {
	return !info.IsDir() && strings.HasSuffix(info.Name(), ".go")
}
