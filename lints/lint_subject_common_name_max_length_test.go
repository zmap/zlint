// lint_subject_common_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectCommonNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectCommonNameLengthGood.pem"
	expected := Pass
	out := Lints["e_subject_common_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectCommonNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectCommonNameLong.pem"
	expected := Error
	out := Lints["e_subject_common_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
