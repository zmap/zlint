// lint_ext_san_dns_name_too_long_test.go
package lints

import (
	"testing"
)

func TestSANDNSShort(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["e_ext_san_dns_name_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANDNSTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSTooLong.pem"
	desEnum := Error
	out := Lints["e_ext_san_dns_name_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
