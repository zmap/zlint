// lint_subject_org_without_country_test.go
package lints

import (
	"testing"
)

func TestOrgNoCoun(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoCountry.pem"
	desEnum := Error
	out := Lints["e_subject_org_without_country"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestOrgYesCoun(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["e_subject_org_without_country"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
