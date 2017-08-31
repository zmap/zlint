// lint_sub_ca_aia_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaAiaMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAMissing.pem"
	expected := Error
	out := Lints["e_sub_ca_aia_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSubCaAiaPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAValid.pem"
	expected := Pass
	out := Lints["e_sub_ca_aia_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
