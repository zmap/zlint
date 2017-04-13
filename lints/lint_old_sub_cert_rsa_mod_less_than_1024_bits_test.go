// lint_old_sub_cert_rsa_mod_less_than_1024_bits_test.go
package lints

import (
	"testing"
)

func TestOldSubCertRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubTooSmall.pem"
	desEnum := Error
	out, _ := Lints["e_old_sub_cert_rsa_mod_less_than_1024_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestOldSubCertRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubSmall.pem"
	desEnum := Pass
	out, _ := Lints["e_old_sub_cert_rsa_mod_less_than_1024_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
