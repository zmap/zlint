// lint_validity_time_not_positive_test.go
package lints

import (
	"testing"
)

func TestValidityNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/validityNegative.pem"
	expected := Error
	out := Lints["e_validity_time_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestValidityPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_validity_time_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
