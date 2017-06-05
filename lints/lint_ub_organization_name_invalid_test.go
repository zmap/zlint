// lint_ub_organization_name_invalid_test.go
package lints

import (
	"testing"
)

func TestUbOrganizationNameGood(t *testing.T) {
	inputPath := "../testlint/testCerts/ubOrganizationNameGood.pem"
	desEnum := Pass
	out, _ := Lints["e_ub_organization_name_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUbOrganzationNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/ubOrganizationNameLong.pem"
	desEnum := Error
	out, _ := Lints["e_ub_organization_name_invalid"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
