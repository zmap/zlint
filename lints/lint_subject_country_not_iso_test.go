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
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCountryIsIso(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectValidCountry.pem"
	expected := Pass
	out := Lints["e_subject_country_not_iso"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
