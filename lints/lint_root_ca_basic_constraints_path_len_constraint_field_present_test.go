package lints

import (
	"testing"
)

func TestRootCaMaxLenPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCaMaxPathLenPresent.pem"
	expected := Warn
	out := Lints["w_root_ca_basic_constraints_path_len_constraint_field_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRootCaMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCaMaxPathLenMissing.pem"
	expected := Pass
	out := Lints["w_root_ca_basic_constraints_path_len_constraint_field_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
