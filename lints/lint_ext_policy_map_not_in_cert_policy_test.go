// lint_ext_policy_map_not_in_cert_policy_test.go
package lints

import (
	"testing"
)

func TestPolicyMapInCertPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapIssuerNotInCertPolicy.pem"
	expected := Warn
	out := Lints["w_ext_policy_map_not_in_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPolicyMapNotInCertPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapGood.pem"
	expected := Pass
	out := Lints["w_ext_policy_map_not_in_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


