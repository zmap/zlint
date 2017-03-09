// lint_sub_ca_no_dns_name_contstraints_test.go
package lints

import (
	"testing"
)

func TestNonEmptyPermittedDns(t *testing.T) {
	inputPath := "../testlint/testCerts/nonEmptyPermittedDns.cer"
	desEnum := Pass
	out, _ := Lints["e_sub_ca_no_dns_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExcludeNonEmptyDns(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedDnsBadExcludedDns.cer"
	desEnum := Error
	out, _ := Lints["e_sub_ca_no_dns_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExcludeEmptyDns(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedDnsGoodExcludedDns.cer"
	desEnum := Pass
	out, _ := Lints["e_sub_ca_no_dns_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
