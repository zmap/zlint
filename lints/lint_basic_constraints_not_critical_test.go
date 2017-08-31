// lint_basic_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestBasicConstNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstNotCrit.pem"
	expected := Error
	out := Lints["e_basic_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestBasicConstCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstCrit.pem"
	expected := Pass
	out := Lints["e_basic_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


