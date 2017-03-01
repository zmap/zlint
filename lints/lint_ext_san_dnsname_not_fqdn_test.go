// lint_ext_san_dnsname_not_fqdn_test.go
package lints

import (
	"testing"
)

func TestDnsFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsFqdn.cer"
	desEnum := Pass
	out, _ := Lints["ext_san_dnsname_not_fqdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDnsNotFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/sanOtherName.cer"
	desEnum := Error
	out, _ := Lints["ext_san_dnsname_not_fqdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
