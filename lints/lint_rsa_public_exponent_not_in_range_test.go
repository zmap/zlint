// lint_rsa_public_exponent_not_in_range_test.go
package lints

import (
	"testing"
)

func TestRsaExpNotInRange(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExp.cer"
	desEnum := Warn
	out, _ := Lints["rsa_public_exponent_not_in_range"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaExpInRange(t *testing.T) {
	inputPath := "../testlint/testCerts/validRsaExpRange.cer"
	desEnum := Pass
	out, _ := Lints["rsa_public_exponent_not_in_range"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
