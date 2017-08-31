// lint_ext_san_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestSANInvalidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithInvalidEmail.pem"
	desEnum := Error
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANInvalidEmail2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithInvalidEmail2.pem"
	desEnum := Error
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithValidEmail.pem"
	desEnum := Pass
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
