// lint_subject_locality_without_org_test.go
package lints

import (
	"testing"
)

func TestLocalNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/localNoOrg.pem"
	desEnum := Error
	out, _ := Lints["e_subject_locality_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestLocalYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/localYesOrg.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_locality_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
