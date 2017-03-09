// lint_sub_cert_key_usage_cert_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestCertSignBitSet(t *testing.T) {
	inputPath := "../testlint/testCerts/subKeyUsageInvalid.cer"
	desEnum := Error
	out, _ := Lints["e_sub_cert_key_usage_cert_sign_bit_set"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertSignBitNotSet(t *testing.T) {
	inputPath := "../testlint/testCerts/subKeyUsageValid.cer"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_key_usage_cert_sign_bit_set"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
