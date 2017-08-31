package lints

import "testing"

func TestDNSNameHyphenBeginningSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameHyphenBeginningSLD.pem"
	desEnum := Error
	out := Lints["e_dnsname_hyphen_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameHyphenEndingSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameHyphenEndingSLD.pem"
	desEnum := Error
	out := Lints["e_dnsname_hyphen_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameNoHyphenInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameWildcardCorrect.pem"
	desEnum := Pass
	out := Lints["e_dnsname_hyphen_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
