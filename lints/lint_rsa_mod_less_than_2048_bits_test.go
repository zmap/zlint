// lint_rsa_mod_less_than_2048_bits_test.go
package lints

import (
	"testing"
)

func TestRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/noRsaLength.pem"
	expected := Error
	out := Lints["e_rsa_mod_less_than_2048_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/yesRsaLength.pem"
	expected := Pass
	out := Lints["e_rsa_mod_less_than_2048_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
