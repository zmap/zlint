// lint_subject_postal_without_org_test.go
package lints

import (
	"testing"
)

func TestPostalNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/postalNoOrg.pem"
	expected := Pass
	out := Lints["e_subject_postal_without_org"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPostalYesOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/postalYesOrg.pem"
	expected := Error
	out := Lints["e_subject_postal_without_org"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


