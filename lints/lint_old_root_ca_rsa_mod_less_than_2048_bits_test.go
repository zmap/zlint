// lint_old_root_ca_rsa_mod_less_than_2048_bits_test.go
package lints

import (
	"testing"
)

func TestOldRootRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldRootModTooSmall.pem"
	desEnum := Error
	out := Lints["e_old_root_ca_rsa_mod_less_than_2048_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestOldRootRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldRootModSmall.pem"
	desEnum := Pass
	out := Lints["e_old_root_ca_rsa_mod_less_than_2048_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
