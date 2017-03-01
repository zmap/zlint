// lint_ext_san_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestSanInvalidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/sanWithInvalidEmail.cer"
	desEnum := Error
	out, _ := Lints["ext_san_rfc822_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanInvalidEmail2(t *testing.T) {
	inputPath := "../testlint/testCerts/sanWithInvalidEmail2.cer"
	desEnum := Error
	out, _ := Lints["ext_san_rfc822_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/sanWithValidEmail.cer"
	desEnum := Pass
	out, _ := Lints["ext_san_rfc822_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
