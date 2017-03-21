// lint_ext_san_critical_with_subject_dn_test.go
package lints

import (
	"testing"
)

func TestSANCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCriticalSubjectUncommonOnly.cer"
	desEnum := Warn
	out, _ := Lints["w_ext_san_critical_with_subject_dn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANNotCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["w_ext_san_critical_with_subject_dn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
