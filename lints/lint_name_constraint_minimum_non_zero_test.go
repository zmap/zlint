// lint_name_constraint_minimum_non_zero_test.go
package lints

import (
	"testing"
)

func TestNcMinZero(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	expected := Pass
	out := Lints["e_name_constraint_minimum_non_zero"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNcMinNotZero(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinPres.pem"
	expected := Error
	out := Lints["e_name_constraint_minimum_non_zero"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
