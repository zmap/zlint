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

// As a note, these certificates were not built, but instead grabbed from censys.io/query
// using the following query to find the raw data and match it to validity period
// select raw, parsed.validity.start from certificates.pemtificates where parsed.signature_algorithm.oid = "1.2.840.113549.1.1.5" limit 200

func TestSubCertProvinceProhibited(t *testing.T) {
	inputPath := "subCertProvinceProhibited.pem"
	expected := lint.Error
	out := test.TestLint("e_sub_cert_province_must_appear", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertProvinceNotProhibited(t *testing.T) {
	inputPath := "subCertProvinceNotProhibited.pem"
	expected := lint.Pass
	out := test.TestLint("e_sub_cert_province_must_appear", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
