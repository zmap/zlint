// lint_rsa_public_exponent_too_small_test.go
package lints

import (
	"testing"
)

func TestRsaExpTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExpLength.pem"
	expected := Error
	out := Lints["e_rsa_public_exponent_too_small"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRsaExpNotTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExpLength.pem"
	expected := Pass
	out := Lints["e_rsa_public_exponent_too_small"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


