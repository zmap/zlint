package lints

import (
	"testing"
)

func TestLeftLabelWildcardCorrect(t *testing.T) {
	
	inputPath := "../testlint/testCerts/dnsNameWildcardCorrect.pem"
	expected := Pass
	out := Lints["e_dnsname_left_label_wildcard_correct"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestLeftLabelWildcardIncorrect(t *testing.T) {
	
	inputPath := "../testlint/testCerts/dnsNameWildcardIncorrect.pem"
	expected := Error
	out := Lints["e_dnsname_left_label_wildcard_correct"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
