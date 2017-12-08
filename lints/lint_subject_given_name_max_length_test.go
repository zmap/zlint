// lint_subject_given_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectGivenNameLengthOK(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectGivenName.pem"
	expected := Pass
	out := Lints["e_subject_given_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectGivenNameTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectGivenNameToolLong.pem"
	expected := Error
	out := Lints["e_subject_given_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
