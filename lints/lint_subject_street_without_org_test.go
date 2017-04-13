// lint_subject_street_without_org_test.go
package lints

import (
	"testing"
)

func TestStreetNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/streetNoOrg.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_street_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestStreetYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/streetYesOrg.pem"
	desEnum := Error
	out, _ := Lints["e_subject_street_without_org"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
