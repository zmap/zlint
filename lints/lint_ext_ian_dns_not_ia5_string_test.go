// lint_ext_ian_dns_not_ia5_string_test.go
package lints

import (
	"testing"
)

func TestIANDNSIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSIA5String.pem"
	expected := Pass
	out := Lints["e_ext_ian_dns_not_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANDNSNotIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSNotIA5String.pem"
	expected := Error
	out := Lints["e_ext_ian_dns_not_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
