// lint_sub_cert_crl_distribution_points_marked_critical_test.go
package lints

import (
	"testing"
)

func TestCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistCrit.pem"
	desEnum := Error
	out := Lints["e_sub_cert_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistNoCrit.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_crl_distribution_points_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
