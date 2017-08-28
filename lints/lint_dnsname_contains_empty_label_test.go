package lints

import (
	"testing"
)

func TestDNSNameEmptyLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameEmptyLabel.pem"
	desEnum := Error
	out, _ := Lints["e_dnsname_empty_label"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameNotEmptyLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNotEmptyLabel.pem"
	desEnum := Pass
	out, _ := Lints["e_dnsname_empty_label"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
