// lint_issuer_dn_leading_whitespace_test.go

package lints

import (
	"testing"
)

func TestIssuerDNLeadingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerDNLeadingSpace.cer"
	desEnum := Warn 
	out, _ := Lints["w_issuer_dn_leading_whitespace"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIssuerDNGood(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.cer"
	desEnum := Pass 
	out, _ := Lints["w_issuer_dn_leading_whitespace"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

