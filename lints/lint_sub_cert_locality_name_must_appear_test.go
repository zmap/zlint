// lint_sub_cert_key_usage_crl_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestSubCertLocalityNameMustAppear(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertLocalityNameMustAppear.pem"
	expected := Error
	out := Lints["e_sub_cert_locality_name_must_appear"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertLocalityNameDoesNotNeedToAppear(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertLocalityNameDoesNotNeedToAppear.pem"
	expected := Pass
	out := Lints["e_sub_cert_locality_name_must_appear"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
