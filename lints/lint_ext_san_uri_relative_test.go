//lint_ext_san_uri_relative_test.go
package lints

import (
	"testing"
)

func TestSANURIRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIRelative.pem"
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

func TestSANURIAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIAbsolute.pem"
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
