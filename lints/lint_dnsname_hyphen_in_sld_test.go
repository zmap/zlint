package lints

import "testing"

func TestDNSNameHyphenBeginningSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameHyphenBeginningSLD.pem"
	desEnum := Error
	out, _ := Lints["e_dnsname_hyphen_in_sld"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameHyphenEndingSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameHyphenBeginningSLD.pem"
	desEnum := Error
	out, _ := Lints["e_dnsname_hyphen_in_sld"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameNoHyphenInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/DNSFQDN.pem"
	desEnum := Pass
	out, _ := Lints["e_dnsname_hyphen_in_sld"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
