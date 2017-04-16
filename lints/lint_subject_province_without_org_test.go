// lint_subject_province_without_org_test.go
package lints

import (
	"testing"
)

func TestProvinceNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/provNoOrg.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_province_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestProvinceYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/provYesOrg.pem"
	desEnum := Error
	out, _ := Lints["e_subject_province_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
