package lints

import (
	"testing"
)

func TestKeyCertSignNotCA(t *testing.T) {
	inputPath := "../testlint/testCerts/keyCertSignNotCA.pem"
	expected := Error
	out := Lints["e_ca_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestKeyCertSignCA(t *testing.T) {
	inputPath := "../testlint/testCerts/keyCertSignCA.pem"
	expected := Pass
	out := Lints["e_ca_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
