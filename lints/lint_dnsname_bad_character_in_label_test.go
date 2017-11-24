package lints

import (
	"testing"
)

func TestBadCharacterInDNSLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameBadCharacterInLabel.pem"
	expected := Error
	out := Lints["e_dnsname_bad_character_in_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestClientDNSCertificate(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameClientCert.pem"
	expected := NA
	out := Lints["e_dnsname_bad_character_in_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestClientValidCertificate(t *testing.T) {
	inputPath := "../testlint/testCerts/validComodo.pem"
	expected := Pass
	out := Lints["e_dnsname_bad_character_in_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
