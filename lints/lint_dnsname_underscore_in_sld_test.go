package lints

import "testing"

func TestDNSNameUnderscoreInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameUnderscoreInSLD.pem"
	expected := Error
	out := Lints["e_dnsname_underscore_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameNoUnderscoreInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameNoUnderscoreInSLD.pem"
	expected := Pass
	out := Lints["e_dnsname_underscore_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


