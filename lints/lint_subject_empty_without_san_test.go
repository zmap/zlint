// lint_subject_empty_without_san_test.go
package lints

import (
	"testing"
)

func TestSubEmptyNoSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectEmptyNoSAN.cer"
	desEnum := Error
	out, _ := Lints["e_subject_empty_without_san"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubEmptyYesSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANSubjectEmptyNotCritical.cer"
	desEnum := Pass
	out, _ := Lints["e_subject_empty_without_san"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
