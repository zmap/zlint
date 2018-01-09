package lints

import (
	"testing"
)

func TestNoNameConstraint(t *testing.T) {
	inputPath := "../testlint/testCerts/noNameConstraint.pem"
	expected := Error
	out := Lints["e_name_constraint_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestHasNameConstraint(t *testing.T) {
	inputPath := "../testlint/testCerts/yesNameConstraint.pem"
	expected := Pass
	out := Lints["e_name_constraint_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
