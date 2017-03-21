// lint_ext_ian_rfc822_format_invalid_test.go
package lints

import (
	"testing"
)

func TestIANInvalidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/IANInvalidEmail.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_rfc822_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/IANValidEmail.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_rfc822_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
