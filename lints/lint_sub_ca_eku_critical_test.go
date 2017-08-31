// lint_sub_ca_eku_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaEkuCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuCrit.pem"
	expected := Warn
	out := Lints["w_sub_ca_eku_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCaEkuNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuNoCrit.pem"
	expected := Pass
	out := Lints["w_sub_ca_eku_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
