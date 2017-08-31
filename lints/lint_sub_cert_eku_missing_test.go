// lint_sub_cert_eku_server_auth_client_auth_missing_test.go
package lints

import (
	"testing"
)

func TestEkuMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageMissing.pem"
	desEnum := Error
	out := Lints["e_sub_cert_eku_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEkuPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageServClient.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_eku_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
