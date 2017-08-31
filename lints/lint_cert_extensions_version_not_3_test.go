// lint_cert_extensions_verson_not_3_test.go
package lints

import (
	"testing"
)

func TestExtsV2(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certVersion2WithExtension.pem"
	expected := Error
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExtsV3(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caBasicConstCrit.pem"
	expected := Pass
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNoExtsV2(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certVersion2NoExtensions.pem"
	expected := Pass
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
