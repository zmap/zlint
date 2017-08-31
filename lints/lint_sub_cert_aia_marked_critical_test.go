// lint_sub_cert_aia_missing_test.go
package lints

import (
	"testing"
)

func TestSubCertAiaMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertAIAMarkedCritical.pem"
	expected := Error
	out := Lints["e_sub_cert_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertAiaNotMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertAIANotMarkedCritical.pem"
	expected := Pass
	out := Lints["e_sub_cert_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
