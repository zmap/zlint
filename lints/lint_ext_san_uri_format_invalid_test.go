//lint_ext_san_uri_format_invalid
package lints

import (
	"testing"
)

func TestSANURIValid(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uri_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANURINoScheme(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINoScheme.pem"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANURINoSchemeSpecificPart(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINoSchemeSpecificPart.pem"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
