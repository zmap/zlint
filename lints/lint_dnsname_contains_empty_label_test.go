package lints

import (
	"testing"
)

func TestDNSNameEmptyLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameEmptyLabel.pem"
	expected := Error
	out := Lints["e_dnsname_empty_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameNotEmptyLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNotEmptyLabel.pem"
	expected := Pass
	out := Lints["e_dnsname_empty_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


