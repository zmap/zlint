// lint_rsa_public_exponent_not_odd_test.go
package lints

import (
	"testing"
)

func TestRsaExpEven(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExp.pem"
	desEnum := Error
	out := Lints["e_rsa_public_exponent_not_odd"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaExpOdd(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExp.pem"
	desEnum := Pass
	out := Lints["e_rsa_public_exponent_not_odd"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
