// lint_ca_country_name_invalid_test.go
package lints

import (
	"testing"
)

func TestCaCountryNameInvalid(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caInvalCountryCode.pem"
	desEnum := Error
	out := Lints["e_ca_country_name_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCaCountryNameValid(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caValCountry.pem"
	desEnum := Pass
	out := Lints["e_ca_country_name_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
