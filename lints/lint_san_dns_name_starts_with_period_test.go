// lint_san_dns_name_starts_with_period_test.go
package lints

import (
	"testing"
)

func TestBrSANDNSStartsWithPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSPeriod.pem"
	desEnum := Error
	out := Lints["e_san_dns_name_starts_with_period"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrSANDNSNotPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	desEnum := Pass
	out := Lints["e_san_dns_name_starts_with_period"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
