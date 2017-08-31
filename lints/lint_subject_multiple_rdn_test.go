// lint_subject_multiple_attr_in_rdn_test.go

package lints

import (
	"testing"
)

func TestSubjectRDNTwoAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectRDNTwoAttribute.pem"
	desEnum := Warn
	out := Lints["w_multiple_subject_rdn"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectRDNOneAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.pem"
	desEnum := Pass
	out := Lints["w_multiple_subject_rdn"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
