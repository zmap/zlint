// lint_ca_key_usage_missing_test.go
package lints

import (
	"testing"
)

func TestCaKeyUsageMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.pem"
	expected := Error
	out := Lints["e_ca_key_usage_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestKeyUsagePresent(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	expected := Pass
	out := Lints["e_ca_key_usage_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


