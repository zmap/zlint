// lint_ext_crl_distribution_marked_critical_test.go
package lints

import (
	"testing"
)

func TestCRLDistribCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistCrit.pem"
	expected := Warn
	out := Lints["w_ext_crl_distribution_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCRLDistribNoCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistNoCrit.pem"
	expected := Pass
	out := Lints["w_ext_crl_distribution_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
