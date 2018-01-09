package lints

import (
	"testing"
)

func TestNcNoEDI(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	expected := Pass
	out := Lints["w_name_constraint_on_edi_party_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNcEDI(t *testing.T) {
	inputPath := "../testlint/testCerts/ncOnEDI.pem"
	expected := Warn
	out := Lints["w_name_constraint_on_edi_party_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
