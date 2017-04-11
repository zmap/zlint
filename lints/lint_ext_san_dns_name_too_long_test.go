// lint_ext_san_dns_name_too_long_test.go
package lints

import (
	"testing"
)

func TestSANDNSShort(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_dns_name_underscore"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANDNSTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSTooLong.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_dns_name_underscore"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
