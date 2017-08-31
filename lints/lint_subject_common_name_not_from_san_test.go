// lint_subject_common_name_not_from_san_test.go
package lints

import (
	"testing"
)

func TestCnNotFromSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithMissingCN.pem"
	desEnum := Error
	out := Lints["e_subject_common_name_not_from_san"].Execute(ReadCertificate(inputPath))
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
	out := Lints["e_subject_common_name_not_from_san"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
