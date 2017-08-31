package lints

import (
	"testing"
)

func TestBadCharacterInDNSLabel(t *testing.T) {
	
	inputPath := "../testlint/testCerts/dnsNameBadCharacterInLabel.pem"
	expected := Error
	out := Lints["e_dnsname_bad_character_in_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
