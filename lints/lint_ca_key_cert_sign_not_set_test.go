// lint_ca_key_cert_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestCaKeyUsageNoCertSign(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNoCertSign.pem"
	expected := Error
	out := Lints["e_ca_key_cert_sign_not_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestKeyUsageCertSign(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	expected := Pass
	out := Lints["e_ca_key_cert_sign_not_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

