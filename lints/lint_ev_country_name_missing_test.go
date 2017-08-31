package lints

import (
	"testing"
)

func TestEvHasCountry(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	expected := Pass
	out := Lints["e_ev_country_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestEvNoCountry(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoCountry.pem"
	expected := Error
	out := Lints["e_ev_country_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
