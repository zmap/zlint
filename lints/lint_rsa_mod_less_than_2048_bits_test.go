// lint_rsa_mod_less_than_2048_bits_test.go
package lints

import (

	"testing"
)

func TestRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/noRsaLength.cer"
	desEnum := Error
	out, _ := Lints["rsa_mod_less_than_2048_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/yesRsaLength.cer"
	desEnum := Pass
	out, _ := Lints["rsa_mod_less_than_2048_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
