package lints

import "testing"

func TestDNSNameLabelTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameLabelTooLong.pem"
	desEnum := Error
	out := Lints["e_dnsname_label_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
