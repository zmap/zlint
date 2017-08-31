// lint_sub_cert_eku_server_auth_client_auth_missing_test.go
package lints

import (
	"testing"
)

func TestEkuBothPres(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageCodeSign.pem"
	desEnum := Error
	out := Lints["e_sub_cert_eku_server_auth_client_auth_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestEkuNeitherPres(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageServClient.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_eku_server_auth_client_auth_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
