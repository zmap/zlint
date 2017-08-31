//lint_ext_san_uri_relative_test.go
package lints

import (
	"testing"
)

func TestSANURIRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIRelative.pem"
	expected := Error
	out := Lints["e_ext_san_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANURIAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIAbsolute.pem"
	expected := Pass
	out := Lints["e_ext_san_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
