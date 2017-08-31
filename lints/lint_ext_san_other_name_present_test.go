// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSANOtherNamePresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANOtherName.pem"
	expected := Error
	out := Lints["e_ext_san_other_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSANOtherNameMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANEDIParty.pem"
	expected := Pass
	out := Lints["e_ext_san_other_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
