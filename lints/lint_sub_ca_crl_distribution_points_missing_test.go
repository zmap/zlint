// lint_sub_ca_crl_distribution_points_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaCrlMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNocrlDist.pem"
	desEnum := Error
	out := Lints["e_sub_ca_crl_distribution_points_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaCrlPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistNoCrit.pem"
	desEnum := Pass
	out := Lints["e_sub_ca_crl_distribution_points_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
