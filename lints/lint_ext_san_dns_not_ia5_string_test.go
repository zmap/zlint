// lint_ext_san_dns_not_ia5_string_test.go
package lints

import (
	"testing"
)

func TestSANDNSNotIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSNotIA5String.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_dns_not_ia5_string"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANDNSIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_dns_not_ia5_string"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
