// subject_contains_noninformational_value_test.go
package lints

import (
	"testing"
)

func TestSubjectNotInformational(t *testing.T) {
	inputPath := "../testlint/testCerts/illegalChar.pem"
	expected := Error
	out := Lints["e_subject_contains_noninformational_value"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectInformational(t *testing.T) {
	inputPath := "../testlint/testCerts/legalChar.pem"
	expected := Pass
	out := Lints["e_subject_contains_noninformational_value"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
