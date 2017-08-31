package lints

import (
	"testing"
)

func TestDNSNameEmptyLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameEmptyLabel.pem"
	expected := Error
	out := Lints["e_dnsname_empty_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestDNSNameNotEmptyLabel(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNotEmptyLabel.pem"
	expected := Pass
	out := Lints["e_dnsname_empty_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
