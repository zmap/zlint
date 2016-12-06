// lint_ian_dns_name_starts_with_period_test.go
package lints

import (

	"testing"
)

func TestBrIanDnsStartsWithPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/ianDnsPeriod.cer"
	desEnum := Error
	out, _ := Lints["ian_dns_name_starts_with_period"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIanDnsNotPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURIValid.cer"
	desEnum := Pass
	out, _ := Lints["ian_dns_name_starts_with_period"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
