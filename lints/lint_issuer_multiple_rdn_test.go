package lints

import (
	"testing"
)

func TestIssuerRDNTwoAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerRDNTwoAttribute.pem"
	expected := Warn
	out := Lints["w_multiple_issuer_rdn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIssuerRDNOneAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.pem"
	expected := Pass
	out := Lints["w_multiple_issuer_rdn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
