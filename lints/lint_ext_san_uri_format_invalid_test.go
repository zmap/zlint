//lint_ext_san_uri_format_invalid
package lints

import (
	"testing"
)

func TestSanURIValid(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.cer"
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

func TestSanURINoScheme(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINoScheme.cer"
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

func TestSanURINoSchemeSpecificPart(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINoSchemeSpecificPart.cer"
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
