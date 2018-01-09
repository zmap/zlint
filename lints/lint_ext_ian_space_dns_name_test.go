package lints

import (
	"testing"
)

func TestIANEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmptyDNS.pem"
	expected := Error
	out := Lints["e_ext_ian_space_dns_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANNotEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/IANNonEmptyDNS.pem"
	expected := Pass
	out := Lints["e_ext_ian_space_dns_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
