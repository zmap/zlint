// lint_ext_freshest_crl_marked_critical_test.go
package lints

import (
	"testing"
)

func TestFreshestCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/frshCRLCritical.pem"
	desEnum := Error
	out := Lints["e_ext_freshest_crl_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestFreshestCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/frshCRLNotCritical.pem"
	desEnum := Pass
	out := Lints["e_ext_freshest_crl_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
