// lint_rsa_public_exponent_too_small_test.go
package lints

import (
	"testing"
)

func TestRsaExpTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExpLength.cer"
	desEnum := Error
	out, _ := Lints["e_rsa_public_exponent_too_small"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaExpNotTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExpLength.cer"
	desEnum := Pass
	out, _ := Lints["e_rsa_public_exponent_too_small"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
