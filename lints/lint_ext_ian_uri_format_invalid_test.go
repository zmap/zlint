//lint_ext_ian_uri_format_invalid_invalid
package lints

import (
	"testing"
)

func TestIanURIValid(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_uri_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanURINoScheme(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoScheme.cer"
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

func TestIanURINoSchemeSpecificPart(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoSchemeSpecificPart.cer"
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
