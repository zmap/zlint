// lint_subject_info_access_marked_critical_test.go
package lints

import (
	"testing"
)

func TestSiaCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/siaCrit.pem"
	expected := Error
	out := Lints["e_subject_info_access_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSiaNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/siaNotCrit.pem"
	expected := Pass
	out := Lints["e_subject_info_access_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
