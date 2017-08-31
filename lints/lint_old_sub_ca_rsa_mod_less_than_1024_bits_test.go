// lint_old_sub_ca_rsa_mod_less_than_1024_bits_test.go
package lints

import (
	"testing"
)

func TestOldCaRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubModTooSmall.pem"
	desEnum := Error
	out := Lints["e_old_sub_ca_rsa_mod_less_than_1024_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestOldCaRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubModSmall.pem"
	desEnum := Pass
	out := Lints["e_old_sub_ca_rsa_mod_less_than_1024_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
