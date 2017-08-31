package lints

import (
	"testing"
)

func TestIssuerRDNTwoAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerRDNTwoAttribute.pem"
	desEnum := Warn
	out := Lints["w_multiple_issuer_rdn"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIssuerRDNOneAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.pem"
	desEnum := Pass
	out := Lints["w_multiple_issuer_rdn"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
