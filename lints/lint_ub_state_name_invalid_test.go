// lint_ub_state_name_invalid_test.go
package lints

import (
	"testing"
)

func TestUbStateNameGood(t *testing.T) {
	inputPath := "../testlint/testCerts/ubStateNameGood.pem"
	desEnum := Pass
	out, _ := Lints["e_ub_state_name_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUbStateNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/ubStateNameLong.pem"
	desEnum := Error
	out, _ := Lints["e_ub_state_name_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
