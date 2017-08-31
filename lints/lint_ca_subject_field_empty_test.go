// lint_ca_subject_field_empty_test.go
package lints

import (
	"testing"
)

func TestCaSubjectMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caSubjectMissing.pem"
	desEnum := Error
	out := Lints["e_ca_subject_field_empty"].Execute(ReadCertificate(inputPath))
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
	inputPath := "../testlint/testCerts/caValCountry.pem"
	desEnum := Pass
	out := Lints["e_ca_subject_field_empty"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
