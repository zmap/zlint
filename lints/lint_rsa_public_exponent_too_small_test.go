// lint_rsa_public_exponent_too_small_test.go
package lints

import (
	"testing"
)

func TestRsaExpTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExpLength.pem"
	desEnum := Error
	out := Lints["e_rsa_public_exponent_too_small"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaExpNotTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExpLength.pem"
	desEnum := Pass
	out := Lints["e_rsa_public_exponent_too_small"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
