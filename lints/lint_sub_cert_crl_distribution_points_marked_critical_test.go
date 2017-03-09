// lint_sub_cert_crl_distribution_points_marked_critical_test.go
package lints

import (
	"testing"
)

func TestCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistCrit.cer"
	desEnum := Error
	out, _ := Lints["e_sub_cert_crl_distribution_points_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistNoCrit.cer"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_crl_distribution_points_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
