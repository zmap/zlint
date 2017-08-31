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
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/IANValidEmail.pem"
	expected := Pass
	out := Lints["e_ext_ian_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
