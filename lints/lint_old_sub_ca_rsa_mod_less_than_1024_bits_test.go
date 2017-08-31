// lint_old_sub_ca_rsa_mod_less_than_1024_bits_test.go
package lints

import (
	"testing"
)

func TestOldCaRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubModTooSmall.pem"
	expected := Error
	out := Lints["e_old_sub_ca_rsa_mod_less_than_1024_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestOldCaRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubModSmall.pem"
	expected := Pass
	out := Lints["e_old_sub_ca_rsa_mod_less_than_1024_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
