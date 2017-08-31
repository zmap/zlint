// lint_san_dns_name_includes_null_char_test.go
package lints

import (
	"testing"
)

func TestBrSANDNSNull(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSNull.pem"
	desEnum := Error
	out := Lints["e_san_dns_name_includes_null_char"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestBrSANDNSNotNull(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	desEnum := Pass
	out := Lints["e_san_dns_name_includes_null_char"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
