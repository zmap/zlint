// lint_name_constraint_minimum_non_zero_test.go
package lints

import (
	"testing"
)

func TestNcMinZero(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	desEnum := Pass
	out := Lints["e_name_constraint_minimum_non_zero"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNcMinNotZero(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinPres.pem"
	desEnum := Error
	out := Lints["e_name_constraint_minimum_non_zero"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
