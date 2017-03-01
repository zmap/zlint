// lint_ext_san_critical_with_subject_dn_test.go
package lints

import (
	"testing"
)

func TestSanCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/sanCriticalSubjectUncommonOnly.cer"
	desEnum := Warn
	out, _ := Lints["ext_san_critical_with_subject_dn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanNotCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["ext_san_critical_with_subject_dn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
