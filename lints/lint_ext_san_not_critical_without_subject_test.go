// lint_ext_san_not_critical_without_subject_test.go
package lints

import (
	"testing"
)

func TestSubjectEmptySANNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/SANSubjectEmptyNotCritical.pem"
	desEnum := Error
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectEmptySANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCaEmptySubject.pem"
	desEnum := Pass
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectNotEmptySANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCriticalSubjectUncommonOnly.pem"
	desEnum := Pass
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
