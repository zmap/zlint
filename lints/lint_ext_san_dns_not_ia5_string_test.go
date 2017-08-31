// lint_ext_san_dns_not_ia5_string_test.go
package lints

import (
	"testing"
)

func TestSANDNSNotIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSNotIA5String.pem"
	expected := Error
	out := Lints["e_ext_san_dns_not_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANDNSIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	expected := Pass
	out := Lints["e_ext_san_dns_not_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
