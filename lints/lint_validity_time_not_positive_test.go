// lint_validity_time_not_positive_test.go
package lints

import (
	"testing"
)

func TestValidityNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/validityNegative.pem"
	desEnum := Error
	out := Lints["e_validity_time_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestValidityPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	desEnum := Pass
	out := Lints["e_validity_time_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
