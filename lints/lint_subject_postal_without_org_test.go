// lint_subject_postal_without_org_test.go
package lints

import (
	"testing"
)

func TestPostalNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/postalNoOrg.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_postal_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPostalYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/postalYesOrg.pem"
	desEnum := Error
	out, _ := Lints["e_subject_postal_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
