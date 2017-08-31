// lint_ext_ian_uri_host_not_fqdn_or_ip_test.go
package lints

import (
	"testing"
)

func TestIANHostURINotFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostNotFQDNOrIP.pem"
	expected := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostURIFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostFQDN.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostURIIP(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostIP.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostWildcardFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostWildcardFQDN.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostWrongWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostWrongWildcard.pem"
	expected := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANHostAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIHostAsterisk.pem"
	expected := Error
	out := Lints["e_ext_ian_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


