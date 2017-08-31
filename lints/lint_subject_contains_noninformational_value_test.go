// subject_contains_noninformational_value_test.go
package lints

import (
	"testing"
)

func TestSubjectNotInformational(t *testing.T) {
	inputPath := "../testlint/testCerts/illegalChar.pem"
	desEnum := Error
	out := Lints["e_subject_contains_noninformational_value"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectInformational(t *testing.T) {
	inputPath := "../testlint/testCerts/legalChar.pem"
	desEnum := Pass
	out := Lints["e_subject_contains_noninformational_value"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
