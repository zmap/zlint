// lint_dsa_has_param_test.go
package lints

import (
	"testing"
)

func TestDSAHasParam(t *testing.T) {
	inputPath := "../testlint/testCerts/DSAHasParam.cer"
	desEnum := Error
	out, _ := Lints["e_dsa_has_param"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDSANoParam(t *testing.T) {
	inputPath := "../testlint/testCerts/DSANoParam.cer"
	desEnum := Pass
	out, _ := Lints["e_dsa_has_param"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

