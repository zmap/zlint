// lint_ian_dns_name_starts_with_period_test.go
package lints

import (
	"testing"
)

func TestBrIANDNSStartsWithPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSPeriod.pem"
	expected := Error
	out := Lints["e_ian_dns_name_starts_with_period"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestBrIANDNSNotPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_ian_dns_name_starts_with_period"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


