package lints

import (
	"testing"
)

func TestCAOrgNameBlank(t *testing.T) {
	inputPath := "../testlint/testCerts/caOrgNameEmpty.pem"
	expected := Error
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCAOrgNameMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caOrgNameMissing.pem"
	expected := Error
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCAOrgNameValid(t *testing.T) {
	inputPath := "../testlint/testCerts/caValOrgName.pem"
	expected := Pass
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
