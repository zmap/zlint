// lint_ext_subject_key_identifier_missing_ca_test.go
package lints

import (
	"testing"
)

func TestSubCaSkiMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCANoSKI.pem"
	desEnum := Error
	out := Lints["e_ext_subject_key_identifier_missing_ca"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaSkiPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/skiNotCriticalCA.pem"
	desEnum := Pass
	out := Lints["e_ext_subject_key_identifier_missing_ca"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
