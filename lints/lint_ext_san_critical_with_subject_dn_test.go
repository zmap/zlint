// lint_ext_san_critical_with_subject_dn_test.go
package lints

import (
	"testing"
)

func TestSANCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCriticalSubjectUncommonOnly.pem"
	desEnum := Warn
	out := Lints["w_ext_san_critical_with_subject_dn"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANNotCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["w_ext_san_critical_with_subject_dn"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
