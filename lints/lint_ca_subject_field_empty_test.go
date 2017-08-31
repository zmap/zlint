// lint_ca_subject_field_empty_test.go
package lints

import (
	"testing"
)

func TestCaSubjectMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caSubjectMissing.pem"
	expected := Error
	out := Lints["e_ca_subject_field_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaSubjectValid(t *testing.T) {
	inputPath := "../testlint/testCerts/caValCountry.pem"
	expected := Pass
	out := Lints["e_ca_subject_field_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


