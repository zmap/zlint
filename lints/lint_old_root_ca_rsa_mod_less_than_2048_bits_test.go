// lint_old_root_ca_rsa_mod_less_than_2048_bits_test.go
package lints

import (

	"testing"
)

func TestOldRootRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldRootModTooSmall.cer"
	desEnum := Error
	out, _ := Lints["old_root_ca_rsa_mod_less_than_2048_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestOldRootRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldRootModSmall.cer"
	desEnum := Pass
	out, _ := Lints["old_root_ca_rsa_mod_less_than_2048_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
