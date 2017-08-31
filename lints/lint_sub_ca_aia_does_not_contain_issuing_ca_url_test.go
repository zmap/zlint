// lint_sub_ca_aia_does_not_contain_issuing_ca_url_test.go
package lints

import (
	"testing"
)

func TestSubCaAiaNoIssuerUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWOcspURL.pem"
	expected := Warn
	out := Lints["w_sub_ca_aia_does_not_contain_issuing_ca_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCaAiaHasIssuerUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWBothURL.pem"
	expected := Pass
	out := Lints["w_sub_ca_aia_does_not_contain_issuing_ca_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
