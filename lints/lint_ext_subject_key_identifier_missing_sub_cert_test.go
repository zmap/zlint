// lint_ext_subject_key_identifier_missing_sub_cert_test.go
package lints

import (
	"testing"
)

func TestSubCertSkiMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertNoSKI.pem"
	desEnum := Warn
	out := Lints["w_ext_subject_key_identifier_missing_sub_cert"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertSkiPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["w_ext_subject_key_identifier_missing_sub_cert"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
