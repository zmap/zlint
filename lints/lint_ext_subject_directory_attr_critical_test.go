// lint_ext_subject_directory_attr_critical_test.go
package lints

import (
	"testing"
)

func TestSdaCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subDirAttCritical.pem"
	expected := Error
	out := Lints["e_ext_subject_directory_attr_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSdaNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/RFC5280example2.pem"
	expected := Pass
	out := Lints["e_ext_subject_directory_attr_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
