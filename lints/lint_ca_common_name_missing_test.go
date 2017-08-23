// lint_ca_country_name_invalid_test.go
package lints

import (
	"testing"
)

func TestCaCommonNameMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caCommonNameMissing.pem"
	desEnum := Error
	out, _ := Lints["e_ca_common_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaCommonNameNotMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caCommonNameNotMissing.pem"
	desEnum := Pass
	out, _ := Lints["e_ca_common_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
