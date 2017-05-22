// lint_subject_common_name_not_from_san_test.go
package lints

import (
	"testing"
)

func TestCnNotFromSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithMissingCN.pem"
	desEnum := Error
	out, _ := Lints["e_subject_common_name_not_from_san"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCnFromSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRegisteredIdBeginning.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_common_name_not_from_san"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCnEmptySAN(t *testing.T) {
	inputPath := "../testlint/testCerts/EmptySAN.cer"
	desEnum := NA
	out, _ := Lints["e_subject_common_name_not_from_san"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath,
			"expected", desEnum,
			"got", out.Result,
		)
	}
}