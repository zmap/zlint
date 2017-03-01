//lint_ext_san_uri_format_invalid
package lints

import (
	"testing"
)

func TestSanUriValid(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURIValid.cer"
	desEnum := Pass
	out, _ := Lints["ext_san_uri_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanUriNoScheme(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURINoScheme.cer"
	desEnum := Error
	out, _ := Lints["ext_san_uri_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanUriNoSchemeSpecificPart(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURINoSchemeSpecificPart.cer"
	desEnum := Error
	out, _ := Lints["ext_san_uri_format_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
