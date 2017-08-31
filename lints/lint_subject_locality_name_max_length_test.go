// lint_subject_locality_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectLocalityNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectLocalityNameLengthGood.pem"
	desEnum := Pass
	out := Lints["e_subject_locality_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectLocalityNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectLocalityNameLong.pem"
	desEnum := Error
	out := Lints["e_subject_locality_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
