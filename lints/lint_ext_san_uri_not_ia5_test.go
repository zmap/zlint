//lint_ext_san_uri_not_ia5_test.go
package lints

import (
	"testing"
)

func TestSANURIIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIIA5.pem"
	expected := Pass
	out := Lints["e_ext_san_uri_not_ia5"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANURINotIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINotIA5.pem"
	expected := Error
	out := Lints["e_ext_san_uri_not_ia5"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
