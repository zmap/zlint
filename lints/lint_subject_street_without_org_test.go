// lint_subject_street_without_org_test.go
package lints

import (
	"testing"
)

func TestStreetNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/streetNoOrg.pem"
	expected := Pass
	out := Lints["e_subject_street_without_org"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestStreetYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/streetYesOrg.pem"
	expected := Error
	out := Lints["e_subject_street_without_org"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
