// lint_sub_cert_eku_extra_values_test.go
package lints

import (
	"testing"
)

func TestEkuExtra(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageServClientEmailCodeSign.pem"
	desEnum := Warn
	out, _ := Lints["w_sub_cert_eku_extra_values"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEkuNoExtra(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageServClientEmail.pem"
	desEnum := Pass
	out, _ := Lints["w_sub_cert_eku_extra_values"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
