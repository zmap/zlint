// lint_sub_ca_crl_distribution_points_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaCrlMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNocrlDist.pem"
	expected := Error
	out := Lints["e_sub_ca_crl_distribution_points_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaCrlPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistNoCrit.pem"
	expected := Pass
	out := Lints["e_sub_ca_crl_distribution_points_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
