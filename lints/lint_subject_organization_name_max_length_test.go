// lint_subject_organization_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectOrganizationNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationNameLengthGood.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_organization_name_max_length"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectOrganzationNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationNameLong.pem"
	desEnum := Error
	out, _ := Lints["e_subject_organization_name_max_length"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
