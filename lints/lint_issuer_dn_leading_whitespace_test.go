// lint_issuer_dn_leading_whitespace_test.go

package lints

import (
	"testing"
)

func TestIssuerDNLeadingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerDNLeadingSpace.pem"
	desEnum := Warn
	out := Lints["w_issuer_dn_leading_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIssuerDNGood(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out := Lints["w_issuer_dn_leading_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
