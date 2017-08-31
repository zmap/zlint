// lint_ca_crl_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestCaKeyUsageNoCRLSign(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNoCRL.pem"
	expected := Error
	out := Lints["e_ca_crl_sign_not_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestKeyUsageCRLSign(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	expected := Pass
	out := Lints["e_ca_crl_sign_not_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


