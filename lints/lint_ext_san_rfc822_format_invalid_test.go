// lint_ext_san_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestSANInvalidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithInvalidEmail.pem"
	expected := Error
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANInvalidEmail2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithInvalidEmail2.pem"
	expected := Error
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithValidEmail.pem"
	expected := Pass
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
