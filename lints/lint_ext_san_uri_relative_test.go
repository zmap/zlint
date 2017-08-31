//lint_ext_san_uri_relative_test.go
package lints

import (
	"testing"
)

func TestSANURIRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIRelative.pem"
	desEnum := Error
	out := Lints["e_ext_san_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANURIAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIAbsolute.pem"
	desEnum := Pass
	out := Lints["e_ext_san_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
