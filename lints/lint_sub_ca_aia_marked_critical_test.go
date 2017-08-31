package lints

import (
	"testing"
)

func TestSubCAAIAMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAMarkedCritical.pem"
	expected := Error
	out := Lints["e_sub_ca_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCAAIANotMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIANotMarkedCritical.pem"
	expected := Pass
	out := Lints["e_sub_ca_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


