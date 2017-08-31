// lint_ext_key_usage_cert_sign_without_ca_test.go
package lints

import (
	"testing"
)

func TestCertSignNoCa(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageCertSignNoBC.pem"
	expected := Error
	out := Lints["e_ext_key_usage_cert_sign_without_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertSignIsCa(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNoCertSign.pem"
	expected := Pass
	out := Lints["e_ext_key_usage_cert_sign_without_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


