package lints

import "testing"

func TestDNSNameUnderscoreInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameUnderscoreInSLD.pem"
	desEnum := Error
	out := Lints["e_dnsname_underscore_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestDNSNameNoUnderscoreInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNoUnderscoreInSLD.pem"
	desEnum := Pass
	out := Lints["e_dnsname_underscore_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
