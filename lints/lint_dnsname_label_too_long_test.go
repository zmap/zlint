package lints

import "testing"

func TestDNSNameLabelTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameLabelTooLong.pem"
	desEnum := Error
	out, _ := Lints["e_dnsname_label_too_long"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
