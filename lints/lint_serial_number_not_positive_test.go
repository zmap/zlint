// lint_serial_number_not_positive_test.go
package lints

import (
	"testing"
)

func TestSnNeagtive(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberNegative.pem"
	expected := Error
	out := Lints["e_serial_number_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSnNotNeagtive(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberValid.pem"
	expected := Pass
	out := Lints["e_serial_number_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
