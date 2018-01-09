package lints

import (
	"testing"
)

func TestSubCaEkuMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUMissing.pem"
	expected := Notice
	out := Lints["n_sub_ca_eku_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCaEkuNotMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuCrit.pem"
	expected := Pass
	out := Lints["n_sub_ca_eku_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
