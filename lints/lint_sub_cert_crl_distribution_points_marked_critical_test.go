// lint_sub_cert_crl_distribution_points_marked_critical_test.go
package lints

import (
	"testing"
)

func TestCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistCrit.pem"
	expected := Error
	out := Lints["e_sub_cert_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistNoCrit.pem"
	expected := Pass
	out := Lints["e_sub_cert_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
