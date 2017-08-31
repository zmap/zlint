// lint_ca_key_cert_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestSubCertIsNotCA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subCertIsNotCA.pem"
	expected := Pass
	out := Lints["e_sub_cert_not_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertIsCA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subCertIsCA.pem"
	expected := Error
	out := Lints["e_sub_cert_not_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
