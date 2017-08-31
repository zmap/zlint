// lint_ext_ian_uri_host_not_fqdn_or_ip_test.go
package lints

import (
	"testing"
)

func TestIANHostURINotFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostNotFQDNOrIP.pem"
	desEnum := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANHostURIFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostFQDN.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANHostURIIP(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostIP.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANHostWildcardFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostWildcardFQDN.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANHostWrongWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostWrongWildcard.pem"
	desEnum := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANHostAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostAsterisk.pem"
	desEnum := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
