//lint_ext_san_uri_relative_test.go
package lints

import (
	"testing"
)

func TestSanUriRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURIRelative.cer"
	desEnum := Error
	out, _ := Lints["ext_san_uri_relative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanUriAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURIAbsolute.cer"
	desEnum := Pass
	out, _ := Lints["ext_san_uri_relative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
