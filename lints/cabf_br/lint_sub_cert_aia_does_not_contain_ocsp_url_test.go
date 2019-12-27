package cabf_br

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

	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/util"
)

func TestSubCertNoIssuerOcsp(t *testing.T) {
	inputPath := "../../testlint/testCerts/subCertWIssuerURL.pem"
	expected := lint.Error
	out := lint.Lints["e_sub_cert_aia_does_not_contain_ocsp_url"].Execute(util.ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertHasIssuerOcsp(t *testing.T) {
	inputPath := "../../testlint/testCerts/subCertWOcspURL.pem"
	expected := lint.Pass
	out := lint.Lints["e_sub_cert_aia_does_not_contain_ocsp_url"].Execute(util.ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
