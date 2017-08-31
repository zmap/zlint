// lint_ext_aia_marked_critical_test.go
package lints

import (
	"testing"
)

func TestAiaCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/aiaCrit.pem"
	expected := Error
	out := Lints["e_ext_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestAiaNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAValid.pem"
	expected := Pass
	out := Lints["e_ext_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


