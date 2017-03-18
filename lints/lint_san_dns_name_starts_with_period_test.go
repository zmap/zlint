// lint_SAN_dns_name_starts_with_period_test.go
package lints

import (
	"testing"
)

func TestBrSanDnsStartsWithPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDnsPeriod.cer"
	desEnum := Error
	out, _ := Lints["e_SAN_dns_name_starts_with_period"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrSanDnsNotPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_SAN_dns_name_starts_with_period"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
