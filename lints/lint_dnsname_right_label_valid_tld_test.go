package lints

import (
	"testing"
)

func TestDNSNameValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameValidTLD.pem"
	expected := Pass
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameNotValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNotValidTLD.pem"
	expected := Error
	out := Lints["e_dnsname_not_valid_tld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
