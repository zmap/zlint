// lint_san_dns_name_includes_null_char_test.go
package lints

import (
	"testing"
)

func TestBrSANDNSNull(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSNull.pem"
	desEnum := Error
	out, _ := Lints["e_san_dns_name_includes_null_char"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrSANDNSNotNull(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	desEnum := Pass
	out, _ := Lints["e_san_dns_name_includes_null_char"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
