// lint_ext_san_not_critical_without_subject_test.go
package lints

import (
	"testing"
)

func TestSubjectEmptySANNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/SANSubjectEmptyNotCritical.pem"
	expected := Error
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectEmptySANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCaEmptySubject.pem"
	expected := Pass
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectNotEmptySANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCriticalSubjectUncommonOnly.pem"
	expected := Pass
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
