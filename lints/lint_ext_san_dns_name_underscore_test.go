// lint_ext_san_dns_name_underscore_test.go
package lints

import (
	"testing"
)

func TestDNSNameUnderscore(t *testing.T) {
	inputPath := "../testlint/testCerts/SANNameUnderscore.cer"
	desEnum := Warn
	out, _ := Lints["w_ext_san_dns_name_underscore"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameGood(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.cer"
	desEnum := Pass
	out, _ := Lints["w_ext_san_dns_name_underscore"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

