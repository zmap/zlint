//lint_ext_ian_uri_relative_test.go
package lints

import (
	"testing"
)

func TestIanUriRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURINoScheme.cer"
	desEnum := Error
	out, _ := Lints["ext_ian_uri_relative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanUriAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURIValid.cer"
	desEnum := Pass
	out, _ := Lints["ext_ian_uri_relative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
