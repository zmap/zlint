// lint_ext_san_critical_with_subject_dn_test.go
package lints

import (
	"testing"
)

func TestSANCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCriticalSubjectUncommonOnly.pem"
	expected := Warn
	out := Lints["w_ext_san_critical_with_subject_dn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANNotCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValGoodAllFields.pem"
	expected := Pass
	out := Lints["w_ext_san_critical_with_subject_dn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
