// lint_subject_street_without_org_test.go
package lints

import (
	"testing"
)

func TestStreetNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/streetNoOrg.pem"
	desEnum := Pass
	out := Lints["e_subject_street_without_org"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestStreetYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/streetYesOrg.pem"
	desEnum := Error
	out := Lints["e_subject_street_without_org"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
