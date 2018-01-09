package lints

import (
	"testing"
)

func TestIANURIRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoScheme.pem"
	expected := Error
	out := Lints["e_ext_ian_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANURIAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
