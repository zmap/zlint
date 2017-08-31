// lint_ext_ian_critical_test.go
package lints

import (
	"testing"
)

func TestIANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/IANCritical.pem"
	expected := Warn
	out := Lints["w_ext_ian_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/IANNotCritical.pem"
	expected := Pass
	out := Lints["w_ext_ian_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
