package lints

import (
	"testing"
)

func TestIDNDnsNameNotNFKC(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNamesNotNFKC.pem"
	expected := Error
	out := Lints["e_international_dns_name_not_nfkc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIDNDnsNameIsNFKC(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNamesNFKC.pem"
	expected := Pass
	out := Lints["e_international_dns_name_not_nfkc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
