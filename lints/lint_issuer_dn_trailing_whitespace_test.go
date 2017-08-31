// lint_issuer_dn_trailing_whitespace_test.go

package lints

import (
	"testing"
)

func TestIssuerDNTrailingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerDNTrailingSpace.pem"
	expected := Warn
	out := Lints["w_issuer_dn_trailing_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestIssuerDNGood2(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["w_issuer_dn_trailing_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
