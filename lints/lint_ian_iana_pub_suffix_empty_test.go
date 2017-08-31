// lint_ian_iana_pub_suffix_empty_test.go
package lints

import (
	"testing"
)

func TestIANBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareSuffix.pem"
	expected := Warn
	out := Lints["w_ian_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANGoodSuffix.pem"
	expected := Pass
	out := Lints["w_ian_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


