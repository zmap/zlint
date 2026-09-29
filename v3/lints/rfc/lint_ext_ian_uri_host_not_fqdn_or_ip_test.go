package rfc

/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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

func TestIANHostURINotFQDN(t *testing.T) {
	inputPath := "IANURIHostNotFQDNOrIP.pem"
	expected := lint.Error
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostURIFQDN(t *testing.T) {
	inputPath := "IANURIHostFQDN.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostURIIP(t *testing.T) {
	inputPath := "IANURIHostIP.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostWildcardFQDN(t *testing.T) {
	inputPath := "IANURIHostWildcardFQDN.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostWrongWildcard(t *testing.T) {
	inputPath := "IANURIHostWrongWildcard.pem"
	expected := lint.Error
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostAsterisk(t *testing.T) {
	inputPath := "IANURIHostAsterisk.pem"
	expected := lint.Error
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANURINoAuthority(t *testing.T) {
	// This certificate has an IAN with URI=sip:alice@sip.uri.com
	// Since this has no authority section, it should be accepted.
	inputPath := "IANURINoAuthority.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANURIHostOpaqueURN(t *testing.T) {
	// This certificate has an IAN with URI=urn:isbn:0451450523
	// Since this has no authority section, it should be accepted.
	inputPath := "IANURIHostOpaqueURN.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANURIHostIPv6(t *testing.T) {
	// This certificate has an IAN with URI=https://[2001:db8::1]/
	// A bracketed IPv6 literal host should be accepted.
	inputPath := "IANURIHostIPv6.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANURIHostURIExampleCom(t *testing.T) {
	// Control case: this certificate has an IAN with URI=https://example.com/,
	// an authority whose host is a plain FQDN.
	inputPath := "IANURIHostURIExampleCom.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_ian_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
