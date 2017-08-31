// lint_rsa_public_exponent_not_in_range_test.go
package lints

import (
	"testing"
)

func TestRsaExpNotInRange(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExp.pem"
	desEnum := Warn
	out := Lints["w_rsa_public_exponent_not_in_range"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRsaExpInRange(t *testing.T) {
	inputPath := "../testlint/testCerts/validRsaExpRange.pem"
	desEnum := Pass
	out := Lints["w_rsa_public_exponent_not_in_range"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
