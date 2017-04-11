// lint_subject_multiple_attr_in_rdn_test.go

package lints

import (
	"testing"
)

func TestSubjectRDNTwoAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectRDNTwoAttribute.cer"
	desEnum := Warn
	out, _ := Lints["w_subject_multiple_attr_in_rdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectRDNOneAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.cer"
	desEnum := Pass
	out, _ := Lints["w_subject_multiple_attr_in_rdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
