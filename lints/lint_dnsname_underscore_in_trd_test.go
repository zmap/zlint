// lint_distribution_point_missing_ldap_or_uri_test.go
package lints

import (
	"testing"
)

func TestDNSNameUnderscoreInTRD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameUnderscoreInTRD.pem"
	expected := Warn
	out := Lints["w_dnsname_underscore_in_trd"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameNoUnderscoreInTRD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNoUnderscoreInTRD.pem"
	expected := Pass
	out := Lints["w_dnsname_underscore_in_trd"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


