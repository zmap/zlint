// lint_ext_san_dns_name_too_long_test.go
package lints

import (
	"testing"
)

func TestSANDNSShort(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_ext_san_dns_name_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANDNSTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSTooLong.pem"
	expected := Error
	out := Lints["e_ext_san_dns_name_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
