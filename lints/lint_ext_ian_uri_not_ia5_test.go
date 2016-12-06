//lint_ext_ian_uri_not_ia5_test.go
package lints

import (

	"testing"
)

func TestIanUriIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURIIa5.cer"
	desEnum := Pass
	out, _ := Lints["ext_ian_uri_not_ia5"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanUriNotIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURINotIA5.cer"
	desEnum := Error
	out, _ := Lints["ext_ian_uri_not_ia5"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
