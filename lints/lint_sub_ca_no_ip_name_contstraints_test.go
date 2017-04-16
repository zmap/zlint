// lint_sub_ca_no_ip_name_contstraints_test.go
package lints

import (
	"testing"
)

func TestBadExclude(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedIPExcludedIPv4.pem"
	desEnum := Error
	out, _ := Lints["e_sub_ca_no_ip_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBadExcludeV6(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedIPExcludedIPv6.pem"
	desEnum := Error
	out, _ := Lints["e_sub_ca_no_ip_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGoodExclude(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedIPExcludedBoth.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_ca_no_ip_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNonEmptyPermitted(t *testing.T) {
	inputPath := "../testlint/testCerts/nonEmptyPermitted.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_ca_no_ip_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
