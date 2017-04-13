// lint_ext_san_dnsname_not_fqdn_test.go
package lints

import (
	"testing"
)

func TestDNSFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/DNSFQDN.pem"
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

func TestDNSWildcardFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSWildcard.pem"
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

func TestDNSNotFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSNameNotFQDN.pem"
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

func TestDNSWrongWildward(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSWrongWildcard.pem"
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

func TestDNSAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDNSAsterisk.pem"
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
