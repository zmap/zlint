// lint_rsa_mod_less_than_2048_bits_test.go
package lints

import (
	"testing"
)

func TestRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/noRsaLength.pem"
	expected := Error
	out := Lints["e_rsa_mod_less_than_2048_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/yesRsaLength.pem"
	expected := Pass
	out := Lints["e_rsa_mod_less_than_2048_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
