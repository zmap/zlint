// lint_ca_subject_field_empty_test.go
package lints

import (
	"testing"
)

func TestCaSubjectMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caSubjectMissing.cer"
	desEnum := Error
	out, _ := Lints["e_ca_subject_field_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaSubjectValid(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caValCountry.cer"
	desEnum := Pass
	out, _ := Lints["e_ca_subject_field_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
