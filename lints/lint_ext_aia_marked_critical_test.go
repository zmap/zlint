// lint_ext_aia_marked_critical_test.go
package lints

import (
	"testing"
)

func TestAiaCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/aiaCrit.pem"
	desEnum := Error
	out := Lints["e_ext_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestAiaNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAValid.pem"
	desEnum := Pass
	out := Lints["e_ext_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
