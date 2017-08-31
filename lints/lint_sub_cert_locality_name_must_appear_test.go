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
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertLocalityNameDoesNotNeedToAppear(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertLocalityNameDoesNotNeedToAppear.pem"
	expected := Pass
	out := Lints["e_sub_cert_locality_name_must_appear"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


