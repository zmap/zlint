// lint_sub_ca_crl_distribution_points_marked_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistCrit.pem"
	desEnum := Error
	out := Lints["e_sub_ca_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistNoCrit.pem"
	desEnum := Pass
	out := Lints["e_sub_ca_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
