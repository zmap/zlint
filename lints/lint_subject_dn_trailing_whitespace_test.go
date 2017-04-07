// lint_subject_dn_trailing_whitespace_test.go

package lints

import (
	"testing"
)

func TestSubjectDNTrailingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectDNTrailingSpace.cer"
	desEnum := Warn 
	out, _ := Lints["w_subject_dn_trailing_whitespace"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectDNGood2(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.cer"
	desEnum := Pass 
	out, _ := Lints["w_subject_dn_trailing_whitespace"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

