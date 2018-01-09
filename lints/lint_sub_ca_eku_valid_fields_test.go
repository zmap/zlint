package lints

import (
	"testing"
)

func TestSubCAEKUValidFields(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUValidFields.pem"
	expected := Pass
	out := Lints["n_sub_ca_eku_not_technically_constrained"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCAEKUNotValidFields(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUNotValidFields.pem"
	expected := NA
	out := Lints["n_sub_ca_eku_not_technically_constrained"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
