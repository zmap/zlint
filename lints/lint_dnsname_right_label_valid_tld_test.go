package lints

import (
	"testing"
)

func TestDNSNameValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameValidTLD.pem"
	desEnum := Pass
	out, _ := Lints["e_dnsname_not_valid_tld"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameNotValidTLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNotValidTLD.pem"
	desEnum := Error
	out, _ := Lints["e_dnsname_not_valid_tld"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
