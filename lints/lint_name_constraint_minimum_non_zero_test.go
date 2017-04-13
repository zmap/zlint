// lint_name_constraint_minimum_non_zero_test.go
package lints

import (
	"testing"
)

func TestNcMinZero(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	desEnum := Pass
	out, _ := Lints["e_name_constraint_minimum_non_zero"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcMinNotZero(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinPres.pem"
	desEnum := Error
	out, _ := Lints["e_name_constraint_minimum_non_zero"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
