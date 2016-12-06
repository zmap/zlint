// lint_sub_ca_crl_distribution_points_marked_critical_test.go
package lints

import (

	"testing"
)

func TestSubCaCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistCrit.cer"
	desEnum := Error
	out, _ := Lints["sub_ca_crl_distribution_points_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWcrlDistNoCrit.cer"
	desEnum := Pass
	out, _ := Lints["sub_ca_crl_distribution_points_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
