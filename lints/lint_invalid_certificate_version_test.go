// lint_invalid_certificate_version_test.go
package lints

import (
	"testing"
)

func TestCertVersion2(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion2WithExtension.pem"
	expected := Error
	out := Lints["e_invalid_certificate_version"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertVersion3(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion3NoExtensions.pem"
	expected := Pass
	out := Lints["e_invalid_certificate_version"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


