// lint_ext_ian_rfc822_format_invalid_test.go
package lints

import (

	"testing"
)

func TestIanInvalidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/ianInvalidEmail.cer"
	desEnum := Error
	out, _ := Lints["ext_ian_rfc822_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/ianValidEmail.cer"
	desEnum := Pass
	out, _ := Lints["ext_ian_rfc822_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
