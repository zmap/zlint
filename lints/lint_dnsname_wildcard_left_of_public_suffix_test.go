package lints

import (
	"testing"
)

func TestWildcardLeftOfPublicSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameWildcardLeftOfPublicSuffix.pem"
	expected := Warn
	out := Lints["w_dnsname_wildcard_left_of_public_suffix"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestWildcardNotLeftOfPublicSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameWildcardNotLeftOfPublicSuffix.pem"
	expected := Pass
	out := Lints["w_dnsname_wildcard_left_of_public_suffix"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


