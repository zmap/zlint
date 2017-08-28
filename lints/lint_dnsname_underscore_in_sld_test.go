package lints

import "testing"

func TestDNSNameUnderscoreInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameUnderscoreInSLD.pem"
	desEnum := Error
	out, _ := Lints["e_dnsname_underscore_in_sld"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameNoUnderscoreInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNoUnderscoreInSLD.pem"
	desEnum := Pass
	out, _ := Lints["e_dnsname_underscore_in_sld"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
