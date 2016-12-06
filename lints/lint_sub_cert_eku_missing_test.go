// lint_sub_cert_eku_server_auth_client_auth_missing_test.go
package lints

import (

	"testing"
)

func TestEkuMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageMissing.cer"
	desEnum := Error
	out, _ := Lints["sub_cert_eku_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEkuPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageServClient.cer"
	desEnum := Pass
	out, _ := Lints["sub_cert_eku_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
