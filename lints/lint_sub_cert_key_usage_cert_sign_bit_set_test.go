// lint_sub_cert_key_usage_cert_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestCertSignBitSet(t *testing.T) {
	inputPath := "../testlint/testCerts/subKeyUsageInvalid.pem"
	desEnum := Error
	out := Lints["e_sub_cert_key_usage_cert_sign_bit_set"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertSignBitNotSet(t *testing.T) {
	inputPath := "../testlint/testCerts/subKeyUsageValid.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_key_usage_cert_sign_bit_set"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
