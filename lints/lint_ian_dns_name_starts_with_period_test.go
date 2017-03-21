// lint_IAN_dns_name_starts_with_period_test.go
package lints

import (
	"testing"
)

func TestBrIanDNSStartsWithPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSPeriod.cer"
	desEnum := Error
	out, _ := Lints["e_IAN_dns_name_starts_with_period"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIanDNSNotPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_IAN_dns_name_starts_with_period"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
