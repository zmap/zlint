// lint_ext_authority_key_identifier_missing_test.go
package lints

import (
	"testing"
)

func TestAKIMissing(t *testing.T) {
	
	inputPath := "../testlint/testCerts/akiMissing.pem"
	expected := Error
	out := Lints["e_ext_authority_key_identifier_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestAKIPresent(t *testing.T) {
	
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_ext_authority_key_identifier_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
