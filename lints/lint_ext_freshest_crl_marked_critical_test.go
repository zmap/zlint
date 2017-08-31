// lint_ext_freshest_crl_marked_critical_test.go
package lints

import (
	"testing"
)

func TestFreshestCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/frshCRLCritical.pem"
	expected := Error
	out := Lints["e_ext_freshest_crl_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestFreshestCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/frshCRLNotCritical.pem"
	expected := Pass
	out := Lints["e_ext_freshest_crl_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
