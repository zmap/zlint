// subject_contains_noninformational_value_test.go
package lints

import (
	"testing"
)

func TestSubjectNotInformational(t *testing.T) {
	inputPath := "../testlint/testCerts/illegalChar.pem"
	desEnum := Error
	out, _ := Lints["e_subject_contains_noninformational_value"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectInformational(t *testing.T) {
	inputPath := "../testlint/testCerts/legalChar.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_contains_noninformational_value"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
