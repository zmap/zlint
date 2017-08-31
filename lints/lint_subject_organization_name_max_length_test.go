// lint_subject_organization_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectOrganizationNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationNameLengthGood.pem"
	desEnum := Pass
	out := Lints["e_subject_organization_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectOrganzationNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationNameLong.pem"
	desEnum := Error
	out := Lints["e_subject_organization_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
