//lint_ext_san_uri_not_ia5_test.go
package lints

import (
	"testing"
)

func TestSANURIIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIIA5.pem"
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

func TestSANURINotIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINotIA5.pem"
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
