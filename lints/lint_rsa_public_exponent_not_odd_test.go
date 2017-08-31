// lint_rsa_public_exponent_not_odd_test.go
package lints

import (
	"testing"
)

func TestRsaExpEven(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExp.pem"
	expected := Error
	out := Lints["e_rsa_public_exponent_not_odd"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRsaExpOdd(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExp.pem"
	expected := Pass
	out := Lints["e_rsa_public_exponent_not_odd"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
