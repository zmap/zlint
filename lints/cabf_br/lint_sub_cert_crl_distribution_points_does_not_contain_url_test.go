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

func TestCrlNoUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistNoURL.pem"
	expected := Error
	out := Lints["e_sub_cert_crl_distribution_points_does_not_contain_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCrlContainsUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistURL.pem"
	expected := Pass
	out := Lints["e_sub_cert_crl_distribution_points_does_not_contain_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCrlContainsUrlInCompoundFullName(t *testing.T) {
	// Re: https://github.com/zmap/zlint/issues/223
	// Previously, we only grabbed the first entry in the fullName of each
	// DistributionPoint, whereas multiple names are allowed (these are
	// interpreted as different names for the same underlying CRL, i.e.
	// providing an LDAP URI and an HTTP URI -- see section 4.2.1.13 of
	// RFC5280).
	inputPath := "../testlint/testCerts/subCrlDistURLInCompoundFullName.pem"
	expected := Pass
	out := Lints["e_sub_cert_crl_distribution_points_does_not_contain_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
