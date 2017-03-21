//lint_ext_ian_uri_not_ia5_test.go
package lints

import (
	"testing"
)

func TestIANURIIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIIa5.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_uri_not_ia5"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANURINotIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINotIa5.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_uri_not_ia5"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
