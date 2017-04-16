// lint_sub_ca_crl_distribution_points_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaCrlMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNocrlDist.pem"
	desEnum := Error
	out, _ := Lints["e_sub_ca_crl_distribution_points_missing"].ExecuteTest(ReadCertificate(inputPath))
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
	out, _ := Lints["e_sub_ca_crl_distribution_points_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
