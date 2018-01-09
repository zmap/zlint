package lints

import (
	"testing"
)

func TestRootCACertPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCAWithCertPolicy.pem"
	expected := Warn
	out := Lints["w_root_ca_contains_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRootCANoCertPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCAValid.pem"
	expected := Pass
	out := Lints["w_root_ca_contains_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
