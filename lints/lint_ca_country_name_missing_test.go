// lint_ca_country_name_missing_test.go
package lints

import (
	"testing"
)

func TestCaCountryNameMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caBlankCountry.cer"
	desEnum := Error
	out, _ := Lints["ca_country_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaCountryNamePresent(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caValCountry.cer"
	desEnum := Pass
	out, _ := Lints["ca_country_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
