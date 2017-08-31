// lint_name_constraint_minimum_non_zero_test.go
package lints

import (
	"testing"
)

func TestNcMinZero(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	expected := Pass
	out := Lints["e_name_constraint_minimum_non_zero"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNcMinNotZero(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinPres.pem"
	expected := Error
	out := Lints["e_name_constraint_minimum_non_zero"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


