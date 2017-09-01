// lint_validity_time_not_positive_test.go
package lints

import (
	"testing"
)

func TestValidityNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/validityNegative.pem"
	expected := Error
	out := Lints["e_validity_time_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestValidityPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_validity_time_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
