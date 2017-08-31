package lints

import (
	"testing"
)

func TestEvHasLocality(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	expected := Pass
	out := Lints["e_ev_locality_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestEvNoLocality(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoLocal.pem"
	expected := Error
	out := Lints["e_ev_locality_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


