package lints

import (
	"testing"
)

func TestBrSANDNSStartsWithPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSPeriod.pem"
	expected := Error
	out := Lints["e_san_dns_name_starts_with_period"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestBrSANDNSNotPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	expected := Pass
	out := Lints["e_san_dns_name_starts_with_period"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
