// lint_subject_common_name_disallowed_test.go
package lints

import (

	"testing"
)

func TestCommonNameInSan(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNameInSan.cer"
	desEnum := Pass
	out, _ := Lints["subject_common_name_disallowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCommonNameNotInSan(t *testing.T) {
	inputPath := "../testlint/testCerts/utcHasSeconds.cer"
	desEnum := Error
	out, _ := Lints["subject_common_name_disallowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
