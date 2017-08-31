// lint_ext_subject_key_identifier_missing_sub_cert_test.go
package lints

import (
	"testing"
)

func TestSubCertSkiMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertNoSKI.pem"
	expected := Warn
	out := Lints["w_ext_subject_key_identifier_missing_sub_cert"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertSkiPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["w_ext_subject_key_identifier_missing_sub_cert"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
