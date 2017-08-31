//lint_ext_ian_uri_not_ia5_test.go
package lints

import (
	"testing"
)

func TestIANURIIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIIA5String.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_uri_not_ia5"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANURINotIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINotIA5String.pem"
	desEnum := Error
	out := Lints["e_ext_ian_uri_not_ia5"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
