// lint_rsa_mod_factors_smaller_than_752_test.go
package lints

import (

	"testing"
)

func TestRsaModFactorTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/evenRsaMod.cer"
	desEnum := Warn
	out, _ := Lints["rsa_mod_factors_smaller_than_752"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaModFactorNotTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExp.cer"
	desEnum := Pass
	out, _ := Lints["rsa_mod_factors_smaller_than_752"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
