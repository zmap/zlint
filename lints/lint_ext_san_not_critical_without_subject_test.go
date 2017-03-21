// lint_ext_san_not_critical_without_subject_test.go
package lints

import (
	"testing"
)

func TestSubjectEmptySANNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/SANSubjectEmptyNotCritical.cer"
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

func TestSubjectEmptySANCrit(t *testing.T) {
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

func TestSubjectNotEmptySANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCriticalSubjectUncommonOnly.cer"
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
