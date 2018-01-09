package lints

import (
	"testing"
)

func TestSubCertAiaMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertAIAMarkedCritical.pem"
	expected := Error
	out := Lints["e_sub_cert_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertAiaNotMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertAIANotMarkedCritical.pem"
	expected := Pass
	out := Lints["e_sub_cert_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
