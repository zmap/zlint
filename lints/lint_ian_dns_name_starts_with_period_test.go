// lint_ian_dns_name_starts_with_period_test.go
package lints

import (
	"testing"
)

func TestBrIANDNSStartsWithPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSPeriod.pem"
	desEnum := Error
	out := Lints["e_ian_dns_name_starts_with_period"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestBrIANDNSNotPeriod(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	desEnum := Pass
	out := Lints["e_ian_dns_name_starts_with_period"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
