//lint_ext_ian_uri_relative_test.go
package lints

import (
	"testing"
)

func TestIANURIRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoScheme.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_uri_relative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANURIAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_uri_relative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
