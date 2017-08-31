// lint_ca_key_cert_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestKeyCertSignNotCA(t *testing.T) {
	
	inputPath := "../testlint/testCerts/keyCertSignNotCA.pem"
	expected := Error
	out := Lints["e_ca_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestKeyCertSignCA(t *testing.T) {
	
	inputPath := "../testlint/testCerts/keyCertSignCA.pem"
	expected := Pass
	out := Lints["e_ca_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
