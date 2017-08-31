// lint_ca_dig_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestCaKeyUsageNoDigSign(t *testing.T) {
	
	inputPath := "../testlint/testCerts/caKeyUsageNoCertSign.pem"
	expected := Notice
	out := Lints["n_ca_digital_signature_not_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestKeyUsageDigSign(t *testing.T) {
	
	inputPath := "../testlint/testCerts/caKeyUsageWDigSign.pem"
	expected := Pass
	out := Lints["n_ca_digital_signature_not_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
