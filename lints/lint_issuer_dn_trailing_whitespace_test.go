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
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIssuerDNGood2(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["w_issuer_dn_trailing_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
