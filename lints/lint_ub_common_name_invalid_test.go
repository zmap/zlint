// lint_ub_common_name_invalid_test.go
package lints

import (
	"testing"
)

func TestUbCommonNameGood(t *testing.T) {
	inputPath := "../testlint/testCerts/ubCommonNameGood.pem"
	desEnum := Pass
	out, _ := Lints["e_ub_common_name_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUbCommonNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/ubCommonNameLong.pem"
	desEnum := Error
	out, _ := Lints["e_ub_common_name_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
