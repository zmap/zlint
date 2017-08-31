package lints

import "testing"

func TestDNSNameHyphenBeginningSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameHyphenBeginningSLD.pem"
	expected := Error
	out := Lints["e_dnsname_hyphen_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameHyphenEndingSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameHyphenEndingSLD.pem"
	expected := Error
	out := Lints["e_dnsname_hyphen_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDNSNameNoHyphenInSLD(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameWildcardCorrect.pem"
	expected := Pass
	out := Lints["e_dnsname_hyphen_in_sld"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


