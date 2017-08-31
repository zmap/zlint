// lint_sub_cert_eku_extra_values_test.go
package lints

import (
	"testing"
)

func TestEkuExtra(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageServClientEmailCodeSign.pem"
	expected := Warn
	out := Lints["w_sub_cert_eku_extra_values"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestEkuNoExtra(t *testing.T) {
	inputPath := "../testlint/testCerts/subExtKeyUsageServClientEmail.pem"
	expected := Pass
	out := Lints["w_sub_cert_eku_extra_values"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
