package lints

import (
	"testing"
)

func TestIDNDnsNameNotNFKC(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNamesNotNFKC.pem"
	desEnum := Error
	out, _ := Lints["e_international_dns_name_not_nfkc"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIDNDnsNameIsNFKC(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNamesNFKC.pem"
	desEnum := Pass
	out, _ := Lints["e_international_dns_name_not_nfkc"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
