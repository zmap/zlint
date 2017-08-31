// lint_subject_common_name_included_test.go
package lints

import (
	"testing"
)

func TestSubCertValidTimeTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertValidTimeTooLong.pem"
	expected := Error
	out := Lints["e_sub_cert_valid_time_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertValidTimeGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertValidTimeGood.pem"
	expected := Pass
	out := Lints["e_sub_cert_valid_time_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
