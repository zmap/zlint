// lint_sub_cert_key_usage_crl_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestSubCertLocalityNameMustAppear(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertLocalityNameMustAppear.pem"
	desEnum := Error
	out := Lints["e_sub_cert_locality_name_must_appear"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertLocalityNameDoesNotNeedToAppear(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertLocalityNameDoesNotNeedToAppear.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_locality_name_must_appear"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
