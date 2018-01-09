package lints

import (
	"testing"
)

func TestNoIssuerField(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerFieldMissing.pem"
	expected := Error
	out := Lints["e_issuer_field_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestHasIssuerField(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerFieldFilled.pem"
	expected := Pass
	out := Lints["e_issuer_field_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
