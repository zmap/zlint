// lint_rsa_mod_factors_smaller_than_752_test.go
package lints

import (
	"testing"
)

func TestRsaModFactorTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/evenRsaMod.pem"
	desEnum := Warn
	out := Lints["w_rsa_mod_factors_smaller_than_752"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaModFactorNotTooSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/goodRsaExp.pem"
	desEnum := Pass
	out := Lints["w_rsa_mod_factors_smaller_than_752"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
