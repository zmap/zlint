package lints

import (
	"testing"
)

func TestSANURIHostNotFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINotFQDN.pem"
	expected := Error
	out := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIHostWildcardFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIHostWildcardFQDN.pem"
	expected := Pass
	out := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIHostWrongWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIHostWrongWildcard.pem"
	expected := Error
	out := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIHostAsterisk(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIHostAsterisk.pem"
	expected := Error
	out := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURIHostFQDN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIHostFQDN.pem"
	expected := Pass
	out := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
