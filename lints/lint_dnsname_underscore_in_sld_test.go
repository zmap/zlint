package lints

import "testing"

func TestDNSNameUnderscoreInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameUnderscoreInSLD.pem"
	expected := Error
	out := Lints["e_dnsname_underscore_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestDNSNameNoUnderscoreInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNoUnderscoreInSLD.pem"
	expected := Pass
	out := Lints["e_dnsname_underscore_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
