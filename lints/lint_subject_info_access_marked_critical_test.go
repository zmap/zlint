// lint_subject_info_access_marked_critical_test.go
package lints

import (
	"testing"
)

func TestSiaCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/siaCrit.pem"
	desEnum := Error
	out := Lints["e_subject_info_access_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSiaNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/siaNotCrit.pem"
	desEnum := Pass
	out := Lints["e_subject_info_access_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
