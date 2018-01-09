package lints

import (
	"testing"
)

func TestSubjectPostalCodeLengthOK(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectPostalCode.pem"
	expected := Pass
	out := Lints["e_subject_postal_code_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectPostalCodeTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectPostalCodeTooLong.pem"
	expected := Error
	out := Lints["e_subject_postal_code_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
