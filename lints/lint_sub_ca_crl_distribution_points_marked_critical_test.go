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
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistNoCrit.pem"
	expected := Pass
	out := Lints["e_sub_ca_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
