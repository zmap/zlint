// lint_ext_san_empty_name_test.go
package lints

import (
	"testing"
)

func TestSANEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/SANEmptyName.pem"
	expected := Error
	out := Lints["e_ext_san_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANNotEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	expected := Pass
	out := Lints["e_ext_san_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


