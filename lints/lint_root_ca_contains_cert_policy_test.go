// lint_root_ca_contains_cert_policy_test.go
package lints

import (
	"testing"
)

func TestRootCACertPolicy(t *testing.T) {
	
	inputPath := "../testlint/testCerts/rootCAWithCertPolicy.pem"
	expected := Warn
	out := Lints["w_root_ca_contains_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestRootCANoCertPolicy(t *testing.T) {
	
	inputPath := "../testlint/testCerts/rootCAValid.pem"
	expected := Pass
	out := Lints["w_root_ca_contains_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
