// lint_subject_org_without_locality_or_province_test.go
package lints

import (

	"testing"
)

func TestOrgNoLoc(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoLocal.cer"
	desEnum := Pass
	out, _ := Lints["subject_org_without_locality_or_province"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestOrgNoProv(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoProv.cer"
	desEnum := Pass
	out, _ := Lints["subject_org_without_locality_or_province"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestOrgNoBoth(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoBoth.cer"
	desEnum := Error
	out, _ := Lints["subject_org_without_locality_or_province"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
