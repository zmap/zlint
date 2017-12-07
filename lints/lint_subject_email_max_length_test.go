// lint_subject_email_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectEmailLengthOK(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectEmailPresent.pem"
	expected := Pass
	out := Lints["e_subject_email_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectEmailTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/SubjectEmailToolLong.pem"
	expected := Error
	out := Lints["e_subject_email_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
