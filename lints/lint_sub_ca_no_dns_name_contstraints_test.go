// lint_sub_ca_no_dns_name_contstraints_test.go
package lints

import (
	"testing"
)

func TestNonEmptyPermittedDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/nonEmptyPermittedDNS.cer"
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

func TestExcludeNonEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedDNSBadExcludedDNS.cer"
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

func TestExcludeEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedDNSGoodExcludedDNS.cer"
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
