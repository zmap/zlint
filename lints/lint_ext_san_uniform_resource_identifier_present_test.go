// lint_ext_san_uniform_resource_identifier_present_test.go
package lints

import (
	"testing"
)

func TestSANURIMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	expected := Pass
	out := Lints["e_ext_san_uniform_resource_identifier_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIBeginning.pem"
	expected := Error
	out := Lints["e_ext_san_uniform_resource_identifier_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIEnd.pem"
	expected := Error
	out := Lints["e_ext_san_uniform_resource_identifier_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


