package rfc

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

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSerialNumberNotPositive(t *testing.T) {
	data := []struct {
		inputPath string
		expected  lint.LintStatus
	}{
		{
			inputPath: "serialNumberNegative.pem",
			expected:  lint.Error,
		},
		{
			inputPath: "serialNumberValid.pem",
			expected:  lint.Pass,
		},
		{
			inputPath: "serialNumberZero.pem",
			expected:  lint.Error,
		},
	}
	for _, d := range data {
		captured := d
		t.Run(d.inputPath, func(t *testing.T) {
			out := test.TestLint("e_serial_number_not_positive", captured.inputPath)
			if out.Status != captured.expected {
				t.Errorf("%s: expected %s, got %s", captured.inputPath, captured.expected, out.Status)
			}
		})
	}
}
