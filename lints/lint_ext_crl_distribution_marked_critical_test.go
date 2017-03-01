// lint_ext_crl_distribution_marked_critical_test.go
package lints

import (
	"testing"
)

func TestCRLDistribCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subCAWcrlDistCrit.cer"
	desEnum := Warn
	out, _ := Lints["ext_crl_distribution_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCRLDistribNoCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subCAWcrlDistNoCrit.cer"
	desEnum := Pass
	out, _ := Lints["ext_crl_distribution_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
