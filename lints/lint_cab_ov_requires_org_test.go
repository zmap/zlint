// lint_cert_policy_requires_org_test.go
package lints

import (
	"testing"
)

func TestCertPolicyOvHasOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cab_ov_requires_org"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyOvNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValNoOrg.pem"
	expected := Error
	out := Lints["e_cab_ov_requires_org"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
