// lint_rsa_exp_negative_test.go
package lints

import (
	"testing"
)

// func TestRsaExpNegative(t *testing.T) {
// 	inputPath := "../testlint/testCerts/rsaExpNegative.pem"
// 	desEnum := Error
// 	out, _ := Lints["rsa_exp_negative"].ExecuteTest(ReadCertificate(inputPath))
// 	if out.Result != desEnum {
// 		t.Error(
// 			"For", inputPath, /* input path*/
// 			"expected", desEnum, /* The enum you expected */
// 			"got", out.Result, /* Actual Result */
// 		)
// 	}
// }

func TestRsaExpPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	desEnum := Pass
	out, _ := Lints["e_rsa_exp_negative"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
