// lint_cert_extensions_verson_not_3_test.go
package lints

import (
	"testing"
)

func TestExtsV2(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion2WithExtension.pem"
	expected := Error
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExtsV3(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstCrit.pem"
	expected := Pass
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNoExtsV2(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion2NoExtensions.pem"
	expected := Pass
	out := Lints["e_cert_extensions_version_not_3"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


