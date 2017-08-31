// lint_ec_improper_curves_test.go
package lints

import (
	"testing"
)

func TestECP224(t *testing.T) {
	inputPath := "../testlint/testCerts/ecdsaP224.pem"
	expected := Error
	out := Lints["e_ec_improper_curves"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestECP256(t *testing.T) {
	inputPath := "../testlint/testCerts/ecdsaP256.pem"
	expected := Pass
	out := Lints["e_ec_improper_curves"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestECP384(t *testing.T) {
	inputPath := "../testlint/testCerts/ecdsaP384.pem"
	expected := Pass
	out := Lints["e_ec_improper_curves"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestECP521(t *testing.T) {
	inputPath := "../testlint/testCerts/ecdsaP521.pem"
	expected := Pass
	out := Lints["e_ec_improper_curves"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


