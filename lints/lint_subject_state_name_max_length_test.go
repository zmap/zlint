// lint_subject_state_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectStateNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectStateNameLengthGood.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_state_name_max_length"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectStateNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectStateNameLong.pem"
	desEnum := Error
	out, _ := Lints["e_subject_state_name_max_length"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
