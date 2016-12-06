// lint_subject_info_access_marked_critical_test.go
package lints

import (

	"testing"
)

func TestSiaCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/siaCrit.cer"
	desEnum := Error
	out, _ := Lints["subject_info_access_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSiaNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/siaNotCrit.cer"
	desEnum := Pass
	out, _ := Lints["subject_info_access_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
