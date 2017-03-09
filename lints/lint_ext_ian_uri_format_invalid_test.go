//lint_ext_ian_uri_format_invalid_invalid
package lints

import (
	"testing"
)

func TestIanUriValid(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURIValid.cer"
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

func TestIanUriNoScheme(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURINoScheme.cer"
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

func TestIanUriNoSchemeSpecificPart(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURINoSchemeSpecificPart.cer"
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
