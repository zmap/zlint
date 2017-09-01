// lint_rsa_exp_negative_test.go
package lints

import (
	"testing"
)

// func TestRsaExpNegative(t *testing.T) {
// 	inputPath := "../testlint/testCerts/rsaExpNegative.pem"
// 	expected := Error
// 	out := Lints["rsa_exp_negative"].ExecuteTest(ReadCertificate(inputPath))
// 	if out.Result != expected {
// 		t.Error(
// 			"For", inputPath,
// 			"expected", expected,
// 			"got", out.Result,
// 		)
// 	}
// }

func TestRsaExpPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_rsa_exp_negative"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
