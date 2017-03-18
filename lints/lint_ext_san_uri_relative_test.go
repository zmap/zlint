//lint_ext_san_uri_relative_test.go
package lints

import (
	"testing"
)

func TestSanURIRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIRelative.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_relative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanURIAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIAbsolute.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uri_relative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
