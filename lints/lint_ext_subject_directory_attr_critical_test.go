// lint_ext_subject_directory_attr_critical_test.go
package lints

import (
	"testing"
)

func TestSdaCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subDirAttCritical.pem"
	desEnum := Error
	out := Lints["e_ext_subject_directory_attr_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSdaNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/RFC5280example2.pem"
	desEnum := Pass
	out := Lints["e_ext_subject_directory_attr_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
