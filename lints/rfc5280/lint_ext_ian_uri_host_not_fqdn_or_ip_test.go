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

func TestIANHostURINotFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostNotFQDNOrIP.pem"
	expected := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostURIFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostFQDN.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostURIIP(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostIP.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostWildcardFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostWildcardFQDN.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostWrongWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostWrongWildcard.pem"
	expected := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostAsterisk.pem"
	expected := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
