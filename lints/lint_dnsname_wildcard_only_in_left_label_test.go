package lints

import (
	"testing"
)

func TestDNSNameWildcardOnlyInLeftLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameWildcardOnlyInLeftLabel.pem"
	expected := Pass
	out := Lints["e_dnsname_wildcard_only_in_left_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameWildcardNotOnlyInLeftLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameWildcardNotOnlyInLeftLabel.pem"
	expected := Error
	out := Lints["e_dnsname_wildcard_only_in_left_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
