//lint_ext_ian_uri_relative_test.go
package lints

import (
	"testing"
)

func TestIANURIRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoScheme.pem"
	desEnum := Error
	out := Lints["e_ext_ian_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANURIAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
