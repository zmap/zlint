// lint_ext_ian_rfc822_format_invalid_test.go
package lints

import (
	"testing"
)

func TestIANInvalidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/IANInvalidEmail.pem"
	expected := Error
	out := Lints["e_ext_ian_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/IANValidEmail.pem"
	expected := Pass
	out := Lints["e_ext_ian_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
