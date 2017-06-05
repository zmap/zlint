// lint_ub_locality_name_invalid_test.go
package lints

import (
	"testing"
)

func TestUbLocalityNameGood(t *testing.T) {
	inputPath := "../testlint/testCerts/ubLocalityNameGood.pem"
	desEnum := Pass
	out, _ := Lints["e_ub_locality_name_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUbLocalityNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/ubLocalityNameLong.pem"
	desEnum := Error
	out, _ := Lints["e_ub_locality_name_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
