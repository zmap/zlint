package lints

import (
	"testing"
)

func TestNcMaxPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/ncAllPres.pem"
	expected := Error
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNcMinPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinPres.pem"
	expected := Pass
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNcEmptyValue(t *testing.T) {
	inputPath := "../testlint/testCerts/ncEmptyValue.pem"
	expected := Pass
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
