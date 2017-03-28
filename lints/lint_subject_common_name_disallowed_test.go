// lint_subject_common_name_disallowed_test.go
package lints

import (
	"testing"
)

func TestCommonNameInSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNameInSAN.cer"
	desEnum := Pass
	out, _ := Lints["e_subject_common_name_disallowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCommonNameNotInSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/utcHasSeconds.cer"
	desEnum := Error
	out, _ := Lints["e_subject_common_name_disallowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEmptySAN(t *testing.T) {
	inputPath := "../testlint/testCerts/EmptySAN.cer"
	desEnum := NA
	out, _ := Lints["e_subject_common_name_disallowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath,
			"expected", desEnum,
			"got", out.Result,
		)
	}
}