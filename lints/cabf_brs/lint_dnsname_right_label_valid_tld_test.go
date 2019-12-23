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

func TestDNSNameValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameValidTLD.pem"
	expected := Pass
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameNotValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNotValidTLD.pem"
	expected := Error
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

// TestDNSNameNotYetValidTLD lints a certificate that was issued for a DNS name
// with a TLD that was not yet delegated at the time the certificate was issued,
// expecting an error.
func TestDNSNameNotYetValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNotYetValidTLD.pem"
	expected := Error
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

// TestDNSNAmeNoLongerValidTLD lints a certificate that was issued for a DNS
// name with a TLD whose delegation was removed from the root DNS at the time
// the certificate was issued, expecting an error.
func TestDNSNameNoLongerValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNoLongerValidTLD.pem"
	expected := Error
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

// TestDNSNameWasValidTLD lints a certificate that was issued for a DNS name
// with a TLD whose delegation was removed from the root DNS, but not until
// after the certificate was issued, expecting no error.
func TestDNSNameWasValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameWasValidTLD.pem"
	expected := Pass
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

// TestDNSNameOnionTLD lints a certificate that was issued for a DNS name with
// a .onion TLD. This ensures the special casing of the .onion gTLD is handled
// correctly and isn't omitted simply because it is not an ICANN/IANA delegated
// TLD.
func TestDNSNameOnionTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameOnionTLD.pem"
	expected := Pass
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameWithIPInCommonName(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameWithIPInCN.pem"
	expected := Pass
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
