// lint_subject_org_without_locality_or_province_test.go
package lints

import (
	"testing"
)

func TestOrgNoLoc(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoLocal.pem"
	expected := Pass
	out := Lints["e_subject_org_without_locality_or_province"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestOrgNoProv(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoProv.pem"
	expected := Pass
	out := Lints["e_subject_org_without_locality_or_province"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestOrgNoBoth(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoBoth.pem"
	expected := Error
	out := Lints["e_subject_org_without_locality_or_province"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
