//lint_ext_san_uri_not_ia5_test.go
package lints

import (
	"testing"
)

func TestSanUriIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURIIA5.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uri_not_ia5"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanUriNotIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURINotIA5.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_not_ia5"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
