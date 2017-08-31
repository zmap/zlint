// lint_subject_empty_without_san_test.go
package lints

import (
	"testing"
)

func TestSubEmptyNoSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectEmptyNoSAN.pem"
	desEnum := Error
	out := Lints["e_subject_empty_without_san"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubEmptyYesSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANSubjectEmptyNotCritical.pem"
	desEnum := Pass
	out := Lints["e_subject_empty_without_san"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
