package lints

import (
	"testing"
)

func TestNcNoX400(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	expected := Pass
	out := Lints["w_name_constraint_on_x400"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNcX400(t *testing.T) {
	inputPath := "../testlint/testCerts/ncOnX400.pem"
	expected := Warn
	out := Lints["w_name_constraint_on_x400"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
