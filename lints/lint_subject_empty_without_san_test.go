// lint_subject_empty_without_san_test.go
package lints

import (
	"testing"
)

func TestSubEmptyNoSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectEmptyNoSAN.pem"
	expected := Error
	out := Lints["e_subject_empty_without_san"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubEmptyYesSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANSubjectEmptyNotCritical.pem"
	expected := Pass
	out := Lints["e_subject_empty_without_san"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
