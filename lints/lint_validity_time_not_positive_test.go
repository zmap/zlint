// lint_validity_time_not_positive_test.go
package lints

import (
	"testing"
)

func TestValidityNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/validityNegative.cer"
	desEnum := Error
	out, _ := Lints["e_validity_time_not_positive"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestValidityPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_validity_time_not_positive"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
