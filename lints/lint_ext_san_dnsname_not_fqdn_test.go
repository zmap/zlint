// lint_ext_san_dnsname_not_fqdn_test.go
package lints

import (
	"testing"
)

func TestDnsFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsFQDN.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_dnsname_not_fqdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDnsWildcardFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDnsWildcard.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_dnsname_not_fqdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDnsNotFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANOtherName.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_dnsname_not_fqdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDnsWrongWildward(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDnsWrongWildcard.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_dnsname_not_fqdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDnsAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDnsAsterisk.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_dnsname_not_fqdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
