// lint_subject_country_not_iso_test.go
package lints

import (
	"testing"
)

func TestCountryNotIso(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectInvalidCountry.pem"
	desEnum := Error
	out := Lints["e_subject_country_not_iso"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCountryIsIso(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectValidCountry.pem"
	desEnum := Pass
	out := Lints["e_subject_country_not_iso"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
