package lints

import (
	"testing"
)

func TestIDNMalformedUnicode(t *testing.T) {
	inputPath := "../testlint/testCerts/idnMalformedUnicode.pem"
	desEnum := Error
	out := Lints["e_international_dns_name_not_unicode"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIDNCorrectUnicode(t *testing.T) {
	inputPath := "../testlint/testCerts/idnCorrectUnicode.pem"
	desEnum := Pass
	out := Lints["e_international_dns_name_not_unicode"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
