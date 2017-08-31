// lint_ext_name_constraints_not_in_ca_test.go
package lints

import (
	"testing"
)

func TestNameConstraintsNotInCa(t *testing.T) {
	inputPath := "../testlint/testCerts/noNameConstraint.pem"
	expected := Error
	out := Lints["e_ext_name_constraints_not_in_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestNameConstraintsInCa(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNameConstCrit.pem"
	expected := Pass
	out := Lints["e_ext_name_constraints_not_in_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
