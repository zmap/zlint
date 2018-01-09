package lints

import (
	"testing"
)

func TestSubjectOrganizationalUnitNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationalUnitNameLengthGood.pem"
	expected := Pass
	out := Lints["e_subject_organizational_unit_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectOrganzationalUnitNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectOrganizationalUnitNameLong.pem"
	expected := Error
	out := Lints["e_subject_organizational_unit_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
