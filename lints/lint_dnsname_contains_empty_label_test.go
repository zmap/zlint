package lints

import (
	"testing"
)

func TestDNSNameEmptyLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameEmptyLabel.pem"
	desEnum := Error
	out := Lints["e_dnsname_empty_label"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestDNSNameNotEmptyLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNotEmptyLabel.pem"
	desEnum := Pass
	out := Lints["e_dnsname_empty_label"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
