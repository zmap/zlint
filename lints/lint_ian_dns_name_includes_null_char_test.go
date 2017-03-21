// lint_IAN_dns_name_includes_null_char_test.go
package lints

import (
	"testing"
)

func TestBrIanDNSNull(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSNull.cer"
	desEnum := Error
	out, _ := Lints["e_IAN_dns_name_includes_null_char"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIanDNSNotNull(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_IAN_dns_name_includes_null_char"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
