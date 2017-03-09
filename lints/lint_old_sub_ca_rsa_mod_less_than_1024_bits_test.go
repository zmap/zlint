// lint_old_sub_ca_rsa_mod_less_than_1024_bits_test.go
package lints

import (
	"testing"
)

func TestOldCaRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubModTooSmall.cer"
	desEnum := Error
	out, _ := Lints["e_old_sub_ca_rsa_mod_less_than_1024_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestOldCaRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubModSmall.cer"
	desEnum := Pass
	out, _ := Lints["e_old_sub_ca_rsa_mod_less_than_1024_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
