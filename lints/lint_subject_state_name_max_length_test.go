// lint_subject_state_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectStateNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectStateNameLengthGood.pem"
	expected := Pass
	out := Lints["e_subject_state_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectStateNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectStateNameLong.pem"
	expected := Error
	out := Lints["e_subject_state_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
