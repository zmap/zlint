// lint_sub_ca_no_ip_name_contstraints_test.go
package lints

import (
	"testing"
)

func TestBadExclude(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedIpExcludedIPv4.cer"
	desEnum := Error
	out, _ := Lints["sub_ca_no_ip_name_contstraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBadExcludeV6(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedIpExcludedIPv6.cer"
	desEnum := Error
	out, _ := Lints["sub_ca_no_ip_name_contstraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGoodExclude(t *testing.T) {
	inputPath := "../testlint/testCerts/emptyPermittedIpExcludedBoth.cer"
	desEnum := Pass
	out, _ := Lints["sub_ca_no_ip_name_contstraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNonEmptyPermitted(t *testing.T) {
	inputPath := "../testlint/testCerts/nonEmptyPermitted.cer"
	desEnum := Pass
	out, _ := Lints["sub_ca_no_ip_name_contstraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
