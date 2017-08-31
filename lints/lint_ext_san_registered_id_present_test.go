// lint_ext_san_registered_id_present_test.go
package lints

import (
	"testing"
)

func TestSANRegIdMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	expected := Pass
	out := Lints["e_ext_san_registered_id_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANRegIdPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRegisteredIdBeginning.pem"
	expected := Error
	out := Lints["e_ext_san_registered_id_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANRegIdPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRegisteredIdEnd.pem"
	expected := Error
	out := Lints["e_ext_san_registered_id_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
