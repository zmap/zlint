package rfc

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
	"testing"

	"github.com/zmap/zlint/v3/test"

	"github.com/zmap/zlint/v3/lint"
)

type testData struct {
	file string
	want lint.LintStatus
}

var tests = []testData{
	{"serialNumberLarge.pem", lint.Error},
	{"serialNumberValid.pem", lint.Pass},
	{"serialNumberLargeDueToSignedMSB.pem", lint.Error},
}

func TestSNSizeLimit(t *testing.T) {
	for _, data := range tests {
		got := test.TestLint("e_serial_number_longer_than_20_octets", data.file).Status
		if got != data.want {
			t.Errorf("%s: expected %s, got %s", data.file, data.want, got)
		}
	}
}
