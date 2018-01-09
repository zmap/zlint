package lints

import (
	"testing"
)

func TestRsaModFactorTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/evenRsaMod.pem"
	expected := Warn
	out := Lints["w_rsa_mod_factors_smaller_than_752"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRsaModFactorNotTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExp.pem"
	expected := Pass
	out := Lints["w_rsa_mod_factors_smaller_than_752"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
