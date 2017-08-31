// lint_root_ca_extended_key_usage_present_test.go
package lints

import (
	"testing"
)

func TestRootCAEKU(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCAWithEKU.pem"
	expected := Error
	out := Lints["e_root_ca_extended_key_usage_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRootCANoEKU(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCAValid.pem"
	expected := Pass
	out := Lints["e_root_ca_extended_key_usage_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
