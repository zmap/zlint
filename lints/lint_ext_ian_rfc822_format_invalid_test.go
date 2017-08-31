// lint_ext_ian_rfc822_format_invalid_test.go
package lints

import (
	"testing"
)

func TestIANInvalidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/IANInvalidEmail.pem"
	desEnum := Error
	out := Lints["e_ext_ian_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/IANValidEmail.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
