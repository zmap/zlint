// lint_ext_ian_uri_host_not_fqdn_or_ip_test.go
package lints

import (
	"testing"
)

func TestIanHostURINotFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostNotFQDNOrIp.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanHostURIFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostFQDN.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanHostURIIp(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostIp.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanHostWildcardFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostWildcardFQDN.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanHostWrongWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostWrongWildcard.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanHostAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostAsterisk.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
