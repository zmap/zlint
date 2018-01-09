package lints

import (
	"testing"
)

func TestBrSANBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANBareWildcard.pem"
	expected := Error
	out := Lints["e_san_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestBrSANNotBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	expected := Pass
	out := Lints["e_san_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
