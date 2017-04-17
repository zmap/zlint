// lint_subject_dn_leading_whitespace_test.go

package lints

import (
	"testing"
)

func TestSubjectDNLeadingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectDNLeadingSpace.pem"
	desEnum := Warn
	out, _ := Lints["w_subject_dn_leading_whitespace"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectDNGood(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out, _ := Lints["w_subject_dn_leading_whitespace"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
