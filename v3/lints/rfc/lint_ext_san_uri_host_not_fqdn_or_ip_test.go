package rfc

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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

func TestSANURIHostNotFQDN(t *testing.T) {
	inputPath := "SANURINotFQDN.pem"
	expected := lint.Error
	out := test.TestLint("e_ext_san_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIHostWildcardFQDN(t *testing.T) {
	inputPath := "SANURIHostWildcardFQDN.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_san_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIHostWrongWildcard(t *testing.T) {
	inputPath := "SANURIHostWrongWildcard.pem"
	expected := lint.Error
	out := test.TestLint("e_ext_san_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIHostAsterisk(t *testing.T) {
	inputPath := "SANURIHostAsterisk.pem"
	expected := lint.Error
	out := test.TestLint("e_ext_san_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIHostFQDN(t *testing.T) {
	inputPath := "SANURIHostFQDN.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_san_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURINoAuthority(t *testing.T) {
	// This certificate has a SAN with URI=sip:alice@sip.uri.com
	// Since this has no authority section, it should be accepted.
	inputPath := "SANURINoAuthority.pem"
	expected := lint.Pass
	out := test.TestLint("e_ext_san_uri_host_not_fqdn_or_ip", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
