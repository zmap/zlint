// lint_rsa_public_exponent_not_in_range_test.go
package lints

import (
	"testing"
)

func TestRsaExpNotInRange(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExp.pem"
	expected := Warn
	out := Lints["w_rsa_public_exponent_not_in_range"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRsaExpInRange(t *testing.T) {
	inputPath := "../testlint/testCerts/validRsaExpRange.pem"
	expected := Pass
	out := Lints["w_rsa_public_exponent_not_in_range"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


