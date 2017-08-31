// lint_sub_cert_aia_missing_test.go
package lints

import (
	"testing"
)

func TestSubCertAiaMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWNoURL.pem"
	expected := Error
	out := Lints["e_sub_cert_aia_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertAiaPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWBothURL.pem"
	expected := Pass
	out := Lints["e_sub_cert_aia_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


