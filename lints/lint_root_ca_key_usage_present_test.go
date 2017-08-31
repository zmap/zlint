package lints

import (
	"testing"
)

func TestRootCAKeyUsagePresent(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCAKeyUsagePresent.pem"
	expected := Pass
	out := Lints["e_root_ca_key_usage_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRootCAKeyUsageMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCAKeyUsageMissing.pem"
	expected := Error
	out := Lints["e_root_ca_key_usage_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


