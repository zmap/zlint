// lint_sub_cert_key_usage_crl_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestCrlSignBitSet(t *testing.T) {
	inputPath := "../testlint/testCerts/subKeyUsageInvalid.pem"
	expected := Error
	out := Lints["e_sub_cert_key_usage_crl_sign_bit_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCrlSignBitNotSet(t *testing.T) {
	inputPath := "../testlint/testCerts/subKeyUsageValid.pem"
	expected := Pass
	out := Lints["e_sub_cert_key_usage_crl_sign_bit_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
