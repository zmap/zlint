// lint_rsa_no_param_test.go
package lints

import (
	"testing"
)

func TestRSANoParam(t *testing.T) {
	inputPath := "../testlint/testCerts/RSANoParam.cer"
	desEnum := Error
	out, _ := Lints["e_rsa_no_param"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRSAHasParam(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA256Good.cer"
	desEnum := Pass
	out, _ := Lints["e_rsa_no_param"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
