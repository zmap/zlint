package lints

import (
	"testing"
)

func TestPolicyMapFromAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapFromAnyPolicy.pem"
	expected := Error
	out := Lints["e_ext_policy_map_any_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPolicyMapToAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapToAnyPolicy.pem"
	expected := Error
	out := Lints["e_ext_policy_map_any_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPolicyMapToNoAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapGood.pem"
	expected := Pass
	out := Lints["e_ext_policy_map_any_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
