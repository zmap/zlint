package lints

import (
	"testing"
)

func TestIDNMalformedUnicode(t *testing.T) {
	inputPath := "../testlint/testCerts/idnMalformedUnicode.pem"
	expected := Error
	out := Lints["e_international_dns_name_not_unicode"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIDNCorrectUnicode(t *testing.T) {
	inputPath := "../testlint/testCerts/idnCorrectUnicode.pem"
	expected := Pass
	out := Lints["e_international_dns_name_not_unicode"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


