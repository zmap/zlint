// lint_san_dns_name_includes_null_char_test.go
package lints

import (
	"testing"
)

func TestBrSANDNSNull(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSNull.pem"
	expected := Error
	out := Lints["e_san_dns_name_includes_null_char"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestBrSANDNSNotNull(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	expected := Pass
	out := Lints["e_san_dns_name_includes_null_char"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
