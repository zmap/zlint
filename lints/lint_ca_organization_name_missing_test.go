// lint_ca_organization_name_missing_test.go
package lints

import (
	"testing"
)

func TestCAOrgNameBlank(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caOrgNameEmpty.pem"
	desEnum := Error
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCAOrgNameMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caOrgNameMissing.pem"
	desEnum := Error
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCAOrgNameValid(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caValOrgName.pem"
	desEnum := Pass
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
