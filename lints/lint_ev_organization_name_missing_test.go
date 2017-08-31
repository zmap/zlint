package lints

import (
	"testing"
)

func TestEvHasOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	expected := Pass
	out := Lints["e_ev_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestEvNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoOrg.pem"
	expected := Error
	out := Lints["e_ev_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
