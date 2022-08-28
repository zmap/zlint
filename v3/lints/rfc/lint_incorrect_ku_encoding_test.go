/*
 * ZLint Copyright 2022 Regents of the University of Michigan
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

package rfc

import (
	"strings"
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestKuIncorrectEncoding(t *testing.T) {
	data := []struct {
		file    string
		want    lint.LintStatus
		details string
	}{
		{
			"incorrect_unused_bits_in_ku_encoding.pem",
			lint.Error,
			"declared to be 5, but it should be 7",
		},
		{
			"keyUsageCertSignEndEntity.pem",
			lint.Pass,
			"",
		},
	}
	for _, d := range data {
		file := d.file
		want := d.want
		details := d.details
		t.Run(file, func(t *testing.T) {
			got := test.TestLint("e_incorrect_ku_encoding", file)
			if got.Status != want {
				t.Errorf("expected %v got %v", want, got)
			}
			if !strings.Contains(got.Details, details) {
				t.Errorf("expected the returned details to contain '%s' but got %s", details, got.Details)
			}
		})
	}
}
