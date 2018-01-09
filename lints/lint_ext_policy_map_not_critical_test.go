package lints

import (
	"testing"
)

func TestPolicyMapNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapNotCritical.pem"
	expected := Warn
	out := Lints["w_ext_policy_map_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPolicyMapCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapGood.pem"
	expected := Pass
	out := Lints["w_ext_policy_map_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
