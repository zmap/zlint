// lint_subject_country_not_iso_test.go
package lints

import (
	"testing"
)

func TestCountryNotIso(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectInvalidCountry.pem"
	expected := Error
	out := Lints["e_subject_country_not_iso"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCountryIsIso(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectValidCountry.pem"
	expected := Pass
	out := Lints["e_subject_country_not_iso"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
