// lint_ca_crl_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestCaKeyUsageNoCRLSign(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caKeyUsageNoCRL.pem"
	desEnum := Error
	out := Lints["e_ca_crl_sign_not_set"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestKeyUsageCRLSign(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	desEnum := Pass
	out := Lints["e_ca_crl_sign_not_set"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
