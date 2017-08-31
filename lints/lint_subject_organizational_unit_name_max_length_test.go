// lint_subject_organizational_unit_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectOrganizationalUnitNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationalUnitNameLengthGood.pem"
	desEnum := Pass
	out := Lints["e_subject_organizational_unit_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectOrganzationalUnitNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationalUnitNameLong.pem"
	desEnum := Error
	out := Lints["e_subject_organizational_unit_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
