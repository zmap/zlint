// lint_subject_org_without_locality_or_province_test.go
package lints

import (
	"testing"
)

func TestOrgNoLoc(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoLocal.pem"
	desEnum := Pass
	out := Lints["e_subject_org_without_locality_or_province"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestOrgNoProv(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoProv.pem"
	desEnum := Pass
	out := Lints["e_subject_org_without_locality_or_province"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestOrgNoBoth(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoBoth.pem"
	desEnum := Error
	out := Lints["e_subject_org_without_locality_or_province"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
