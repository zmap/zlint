//lint_ext_ian_uri_format_invalid_invalid
package lints

import (
	"testing"
)

func TestIANURIValid(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANURINoScheme(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoScheme.pem"
	desEnum := Error
	out := Lints["e_ext_san_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANURINoSchemeSpecificPart(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoSchemeSpecificPart.pem"
	desEnum := Error
	out := Lints["e_ext_san_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
