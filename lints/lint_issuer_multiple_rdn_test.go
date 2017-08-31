package lints

import (
	"testing"
)

func TestIssuerRDNTwoAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerRDNTwoAttribute.pem"
	expected := Warn
	out := Lints["w_multiple_issuer_rdn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIssuerRDNOneAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.pem"
	expected := Pass
	out := Lints["w_multiple_issuer_rdn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
