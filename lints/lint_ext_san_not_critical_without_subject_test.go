// lint_ext_san_not_critical_without_subject_test.go
package lints

import (
	"testing"
)

func TestSubjectEmptySanNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/sanSubjectEmptyNotCritical.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_not_critical_without_subject"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectEmptySanCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCaEmptySubject.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_not_critical_without_subject"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectNotEmptySanCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/sanCriticalSubjectUncommonOnly.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_not_critical_without_subject"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
