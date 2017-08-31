// lint_ca_key_cert_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestSubCertIsNotCA(t *testing.T) {
	
	inputPath := "../testlint/testCerts/subCertIsNotCA.pem"
	expected := Pass
	out := Lints["e_sub_cert_not_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSubCertIsCA(t *testing.T) {
	
	inputPath := "../testlint/testCerts/subCertIsCA.pem"
	expected := Error
	out := Lints["e_sub_cert_not_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
