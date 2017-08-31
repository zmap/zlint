// lint_ext_subject_key_identifier_missing_ca_test.go
package lints

import (
	"testing"
)

func TestSubCaSkiMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCANoSKI.pem"
	expected := Error
	out := Lints["e_ext_subject_key_identifier_missing_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaSkiPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/skiNotCriticalCA.pem"
	expected := Pass
	out := Lints["e_ext_subject_key_identifier_missing_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
