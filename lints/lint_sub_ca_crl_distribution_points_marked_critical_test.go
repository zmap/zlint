// lint_sub_ca_crl_distribution_points_marked_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistCrit.pem"
	expected := Error
	out := Lints["e_sub_ca_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCaCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistNoCrit.pem"
	expected := Pass
	out := Lints["e_sub_ca_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
