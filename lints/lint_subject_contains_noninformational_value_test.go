package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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
)

func TestSubjectInformational(t *testing.T) {
	testCases := []struct {
		name      string
		inputPath string
		result    LintStatus
	}{
		{
			name:      "simple all legal",
			inputPath: "../testlint/testCerts/legalChar.pem",
			result:    Pass,
		},
		{
			name:      "subject with metadata only",
			inputPath: "../testlint/testCerts/illegalChar.pem",
			result:    Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := Lints["e_subject_contains_noninformational_value"].Execute(ReadCertificate(tc.inputPath))
			if out.Status != tc.result {
				t.Errorf("%s: expected %s, got %s", tc.inputPath, tc.result, out.Status)
			}
		})
	}
}

func TestCheckAlphaNumericOrUTF8Present(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "ascii lowercase",
			input:  "aa",
			result: true,
		},
		{
			name:   "ascii uppercase",
			input:  "AA",
			result: true,
		},
		{
			name:   "ascii numbers",
			input:  "123",
			result: true,
		},
		{
			name:   "ascii start with metadata",
			input:  "-- abc3",
			result: true,
		},
		{
			name:   "ascii end with metadata",
			input:  "abc3 ..",
			result: true,
		},
		{
			name:   "UTF8",
			input:  "テスト",
			result: true,
		},
		{
			name:   "UTF8 start with metadata",
			input:  "?? テスト",
			result: true,
		},
		{
			name:   "UTF8 end with metadata",
			input:  "テスト ??",
			result: true,
		},
		{
			name:   "-",
			input:  "-",
			result: false,
		},
		{
			name:   "**",
			input:  "**",
			result: false,
		},
		{
			name:   "...",
			input:  "...",
			result: false,
		},
		{
			name:   "- -",
			input:  "- -",
			result: false,
		},
		{
			name:   " -",
			input:  " -",
			result: false,
		},
		{
			name:   " ",
			input:  " ",
			result: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := checkAlphaNumericOrUTF8Present(tc.input)
			if result != tc.result {
				t.Errorf("expected check to be %v, got %v", tc.result, result)
			}
		})
	}
}
