// lint_subject_province_without_org_test.go
package lints

import (
	"testing"
)

func TestProvinceNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/provNoOrg.cer"
	desEnum := Pass
	out, _ := Lints["subject_province_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestProvinceYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/provYesOrg.cer"
	desEnum := Error
	out, _ := Lints["subject_province_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
