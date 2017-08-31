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
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestCRLDistHttp(t *testing.T) {
	
	inputPath := "../testlint/testCerts/crlDistribWithHTTP.pem"
	expected := Pass
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestCRLDistLdap(t *testing.T) {
	
	inputPath := "../testlint/testCerts/crlDistribWithLDAP.pem"
	expected := Pass
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
