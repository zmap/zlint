// lint_ext_subject_key_identifier_missing_ca_test.go
package lints

import (
	"testing"
)

func TestSubCaSkiMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCANoSKI.cer"
	desEnum := Error
	out, _ := Lints["ext_subject_key_identifier_missing_ca"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaSkiPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/skiNotCriticalCA.cer"
	desEnum := Pass
	out, _ := Lints["ext_subject_key_identifier_missing_ca"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
