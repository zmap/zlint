package cabf_br

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

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestNotApplicable(t *testing.T) {
	inputPath := "tls_cert_20220830_with_OU.pem"
	expected := lint.NA
	out := test.TestLint("e_server_cert_subject_contains_ou", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestOUPresentCase1(t *testing.T) {
	inputPath := "tls_cert_20220901_with_OU.pem"
	expected := lint.Error
	out := test.TestLint("e_server_cert_subject_contains_ou", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestOUPresentCase2(t *testing.T) {
	inputPath := "tls_cert_20221023_with_OU.pem"
	expected := lint.Error
	out := test.TestLint("e_server_cert_subject_contains_ou", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestOUNotPresent(t *testing.T) {
	inputPath := "tls_cert_20220902_wout_OU.pem"
	expected := lint.Pass
	out := test.TestLint("e_server_cert_subject_contains_ou", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
