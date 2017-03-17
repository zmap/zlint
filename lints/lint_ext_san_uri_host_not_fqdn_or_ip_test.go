// lint_ext_san_uri_host_not_fqdn_or_ip_test.go
package lints

import (
	"testing"
)

func TestSanUriHostNotFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/sanUriNotFqdn.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanUriHostWildcardFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/sanUriHostWildcardFqdn.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanUriHostWrongWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/sanUriHostWrongWildcard.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanUriHostAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/sanUriHostAsterisk.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanUriHostFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/sanUriHostFqdn.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}


