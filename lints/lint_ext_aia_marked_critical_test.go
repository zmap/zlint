// lint_ext_aia_marked_critical_test.go
package lints

import (

	"testing"
)

func TestAiaCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/aiaCrit.cer"
	desEnum := Error
	out, _ := Lints["ext_aia_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestAiaNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAValid.cer"
	desEnum := Pass
	out, _ := Lints["ext_aia_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
