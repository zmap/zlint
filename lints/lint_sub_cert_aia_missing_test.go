// lint_sub_cert_aia_missing_test.go
package lints

import (
	"testing"
)

func TestSubCertAiaMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWNoURL.pem"
	expected := Error
	out := Lints["e_sub_cert_aia_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertAiaPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWBothURL.pem"
	expected := Pass
	out := Lints["e_sub_cert_aia_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
