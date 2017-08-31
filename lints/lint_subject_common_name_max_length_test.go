// lint_subject_common_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectCommonNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectCommonNameLengthGood.pem"
	desEnum := Pass
	out := Lints["e_subject_common_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectCommonNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectCommonNameLong.pem"
	desEnum := Error
	out := Lints["e_subject_common_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
