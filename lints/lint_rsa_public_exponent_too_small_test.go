// lint_rsa_public_exponent_too_small_test.go
package lints

import (
	"testing"
)

func TestRsaExpTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/badRsaExpLength.pem"
	expected := Error
	out := Lints["e_rsa_public_exponent_too_small"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRsaExpNotTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExpLength.pem"
	expected := Pass
	out := Lints["e_rsa_public_exponent_too_small"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
