// lint_ext_ian_dns_not_ia5_string_test.go
package lints

import (
	"testing"
)

func TestIANDNSIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSIA5String.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_dns_not_ia5_string"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANDNSNotIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSNotIA5String.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_dns_not_ia5_string"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
