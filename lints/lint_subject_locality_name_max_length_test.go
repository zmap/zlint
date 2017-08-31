// lint_subject_locality_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectLocalityNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectLocalityNameLengthGood.pem"
	expected := Pass
	out := Lints["e_subject_locality_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectLocalityNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectLocalityNameLong.pem"
	expected := Error
	out := Lints["e_subject_locality_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
