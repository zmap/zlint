// lint_subject_common_name_not_from_san_test.go
package lints

import (
	"testing"
)

func TestCnNotFromSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithMissingCN.pem"
	expected := Error
	out := Lints["e_subject_common_name_not_from_san"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCnFromSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRegisteredIdBeginning.pem"
	expected := Pass
	out := Lints["e_subject_common_name_not_from_san"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
