// lint_subject_org_without_country_test.go
package lints

import (
	"testing"
)

func TestOrgNoCoun(t *testing.T) {
	inputPath := "../testlint/testCerts/orgNoCountry.pem"
	expected := Error
	out := Lints["e_subject_org_without_country"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestOrgYesCoun(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_subject_org_without_country"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
