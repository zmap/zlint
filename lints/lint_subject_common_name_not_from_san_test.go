// lint_subject_common_name_not_from_san_test.go
package lints

import (

	"testing"
)

func TestCnNotFromSan(t *testing.T) {
	inputPath := "../testlint/testCerts/sanWithMissingCN.cer"
	desEnum := Error
	out, _ := Lints["subject_common_name_not_from_san"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCnFromSan(t *testing.T) {
	inputPath := "../testlint/testCerts/sanRegisteredIdBeginning.cer"
	desEnum := Pass
	out, _ := Lints["subject_common_name_not_from_san"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
