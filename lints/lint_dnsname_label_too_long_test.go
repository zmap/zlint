package lints

import "testing"

func TestDNSNameLabelTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameLabelTooLong.pem"
	expected := Error
	out := Lints["e_dnsname_label_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


