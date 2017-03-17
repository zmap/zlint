// lint_ext_ian_uri_host_not_fqdn_or_ip_test.go
package lints

import (
	"testing"
)

func TestIanHostUriNotFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/ianUriHostNotFqdnOrIp.cer"
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

func TestIanHostUriFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/ianUriHostFqdn.cer"
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

func TestIanHostUriIp(t *testing.T) {
	inputPath := "../testlint/testCerts/ianUriHostIp.cer"
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

func TestIanHostWildcardFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/ianUriHostWildcardFqdn.cer"
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
	inputPath := "../testlint/testCerts/ianUriHostWrongWildcard.cer"
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
	inputPath := "../testlint/testCerts/ianUriHostAsterisk.cer"
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
