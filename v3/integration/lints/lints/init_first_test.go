package lints

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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
	"testing"

	"github.com/zmap/zlint/v3/integration/lints/lint"
)

func TestInitFirst_Lint(t *testing.T) {
	data := []struct {
		inputFile  string
		expectPass bool
	}{
		{inputFile: "testdata/lint_initializeFirst.go", expectPass: true},
		{inputFile: "testdata/lint_initializeFirst.go", expectPass: true},
		{inputFile: "testdata/lint_initializeNotFirst.go", expectPass: false},
		{inputFile: "testdata/lint_initializeFirstNoFunctions.go", expectPass: false},
	}
	l := &InitFirst{}
	for _, test := range data {
		file := test.inputFile
		want := test.expectPass
		t.Run(file, func(t *testing.T) {
			r, err := lint.RunLintForFile(file, l)
			if err != nil {
				t.Fatal(err)
			}
			if want && r != nil {
				t.Errorf("got unexepcted error result, %s", r)
			} else if !want && r == nil {
				t.Errorf("expected failure but got nothing")
			}
		})
	}
}
