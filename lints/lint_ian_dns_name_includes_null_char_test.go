// lint_ian_dns_name_includes_null_char_test.go
package lints

import (
	"testing"
)

func TestBrIANDNSNull(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSNull.pem"
	desEnum := Error
	out := Lints["e_ian_dns_name_includes_null_char"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestBrIANDNSNotNull(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	desEnum := Pass
	out := Lints["e_ian_dns_name_includes_null_char"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
