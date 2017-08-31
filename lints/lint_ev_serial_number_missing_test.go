package lints

import (
	"testing"
)

func TestEvHasSN(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	expected := Pass
	out := Lints["e_ev_serial_number_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestEvNoSN(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoSN.pem"
	expected := Error
	out := Lints["e_ev_serial_number_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
