// lint_issuer_dn_leading_whitespace_test.go

package lints

import (
	"testing"
)

func TestIssuerDNLeadingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerDNLeadingSpace.pem"
	expected := Warn
	out := Lints["w_issuer_dn_leading_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestIssuerDNGood(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["w_issuer_dn_leading_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
