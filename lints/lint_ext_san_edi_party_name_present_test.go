// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSANEDIPartyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANEDIParty.pem"
	expected := Error
	out := Lints["e_ext_san_edi_party_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANEDIPartyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANOtherName.pem"
	expected := Pass
	out := Lints["e_ext_san_edi_party_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


