// lint_issuer_dn_trailing_whitespace_test.go

package lints

import (
	"testing"
)

func TestIssuerDNTrailingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerDNTrailingSpace.pem"
	desEnum := Warn
	out, _ := Lints["w_issuer_dn_trailing_whitespace"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIssuerDNGood2(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out, _ := Lints["w_issuer_dn_trailing_whitespace"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
