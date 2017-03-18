// lint_subject_common_name_not_from_SAN_test.go
package lints

import (
	"testing"
)

func TestCnNotFromSan(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithMissingCN.cer"
	desEnum := Error
	out, _ := Lints["e_subject_common_name_not_from_SAN"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCnFromSan(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRegisteredIdBeginning.cer"
	desEnum := Pass
	out, _ := Lints["e_subject_common_name_not_from_SAN"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
