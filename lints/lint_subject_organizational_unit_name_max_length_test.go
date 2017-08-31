// lint_subject_organizational_unit_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectOrganizationalUnitNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationalUnitNameLengthGood.pem"
	expected := Pass
	out := Lints["e_subject_organizational_unit_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectOrganzationalUnitNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationalUnitNameLong.pem"
	expected := Error
	out := Lints["e_subject_organizational_unit_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
