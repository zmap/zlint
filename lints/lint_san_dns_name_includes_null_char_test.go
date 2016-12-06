// lint_san_dns_name_includes_null_char_test.go
package lints

import (

	"testing"
)

func TestBrSanDnsNull(t *testing.T) {
	inputPath := "../testlint/testCerts/sanDnsNull.cer"
	desEnum := Error
	out, _ := Lints["san_dns_name_includes_null_char"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrSanDnsNotNull(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURIValid.cer"
	desEnum := Pass
	out, _ := Lints["san_dns_name_includes_null_char"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
