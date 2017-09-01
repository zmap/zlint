// lint_distribution_point_missing_ldap_or_uri_test.go
package lints

import (
	"testing"
)

func TestCRLDistNoHttp(t *testing.T) {
	inputPath := "../testlint/testCerts/crlDistribNoHTTP.pem"
	expected := Warn
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCRLDistHttp(t *testing.T) {
	inputPath := "../testlint/testCerts/crlDistribWithHTTP.pem"
	expected := Pass
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCRLDistLdap(t *testing.T) {
	inputPath := "../testlint/testCerts/crlDistribWithLDAP.pem"
	expected := Pass
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
