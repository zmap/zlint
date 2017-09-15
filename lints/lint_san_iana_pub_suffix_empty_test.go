// lint_san_iana_pub_suffix_empty_test.go
package lints

import (
	"testing"
)

func TestSANBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/SANBareSuffix.pem"
	expected := Warn
	out := Lints["w_san_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANBarePrivatePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/sanPrivatePublicSuffix.pem"
	expected := Pass
	out := Lints["w_san_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/SANGoodSuffix.pem"
	expected := Pass
	out := Lints["w_san_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
