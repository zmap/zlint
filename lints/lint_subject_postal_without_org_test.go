// lint_subject_postal_without_org_test.go
package lints

import (

	"testing"
)

func TestPostalNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/postalNoOrg.cer"
	desEnum := Pass
	out, _ := Lints["subject_postal_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPostalYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/postalYesOrg.cer"
	desEnum := Error
	out, _ := Lints["subject_postal_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
