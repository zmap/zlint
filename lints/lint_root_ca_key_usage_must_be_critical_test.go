package lints

import (
	"testing"
)

func TestRootCAKeyUsageCritical(t *testing.T) {
	
	inputPath := "../testlint/testCerts/rootCAKeyUsagePresent.pem"
	expected := Pass
	out := Lints["e_root_ca_key_usage_must_be_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestRootCAKeyUsageNotCritical(t *testing.T) {
	
	inputPath := "../testlint/testCerts/rootCAKeyUsageNotCritical.pem"
	expected := Error
	out := Lints["e_root_ca_key_usage_must_be_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
