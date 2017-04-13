// lint_subject_org_without_country_test.go
package lints

import (
	"testing"
)

func TestOrgNoCoun(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoCountry.pem"
	desEnum := Error
	out, _ := Lints["e_subject_org_without_country"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestOrgYesCoun(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_org_without_country"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
