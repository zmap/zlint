package lints

import (
	"testing"
)

func TestSANURIRelative(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIRelative.pem"
	expected := Error
	out := Lints["e_ext_san_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIAbsolute(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIAbsolute.pem"
	expected := Pass
	out := Lints["e_ext_san_uri_relative"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
