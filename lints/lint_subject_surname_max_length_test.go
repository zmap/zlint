// lint_subject_surname_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectSurnameLengthOK(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectSurname.pem"
	expected := Pass
	out := Lints["e_subject_surname_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectSurnameTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectSurnameTooLong.pem"
	expected := Error
	out := Lints["e_subject_surname_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
