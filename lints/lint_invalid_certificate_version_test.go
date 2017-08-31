// lint_invalid_certificate_version_test.go
package lints

import (
	"testing"
)

func TestCertVersion2(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion2WithExtension.pem"
	expected := Error
	out := Lints["e_invalid_certificate_version"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertVersion3(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion3NoExtensions.pem"
	expected := Pass
	out := Lints["e_invalid_certificate_version"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
