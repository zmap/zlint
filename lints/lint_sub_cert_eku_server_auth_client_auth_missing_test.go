// lint_sub_cert_eku_server_auth_client_auth_missing_test.go
package lints

import (
	"testing"
)

func TestEkuBothPres(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageCodeSign.pem"
	expected := NA
	out := Lints["e_sub_cert_eku_server_auth_client_auth_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestEkuNeitherPres(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageServClient.pem"
	expected := Pass
	out := Lints["e_sub_cert_eku_server_auth_client_auth_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
