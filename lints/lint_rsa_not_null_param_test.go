// lint_rsa_not_null_param_test.go
package lints

import (
	"testing"
)

func TestRSANotNullParam(t *testing.T) {
	inputPath := "../testlint/testCerts/RSANotNullParam.pem"
	desEnum := Error
	out, _ := Lints["e_rsa_not_null_param"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRSANullParam(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA256Good.pem"
	desEnum := Pass
	out, _ := Lints["e_rsa_not_null_param"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
