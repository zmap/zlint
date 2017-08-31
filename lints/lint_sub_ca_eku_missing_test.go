// lint_sub_ca_eku_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaEkuMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUMissing.pem"
	expected := Error
	out := Lints["e_sub_ca_eku_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSubCaEkuNotMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuCrit.pem"
	expected := Pass
	out := Lints["e_sub_ca_eku_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
