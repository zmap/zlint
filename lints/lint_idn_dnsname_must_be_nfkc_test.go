package lints

import (
	"testing"
)

func TestIDNDnsNameNotNFKC(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNamesNotNFKC.pem"
	desEnum := Error
	out := Lints["e_international_dns_name_not_nfkc"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIDNDnsNameIsNFKC(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNamesNFKC.pem"
	desEnum := Pass
	out := Lints["e_international_dns_name_not_nfkc"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
