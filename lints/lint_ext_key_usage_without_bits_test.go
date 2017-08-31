// lint_ext_key_usage_without_bits_test.go
package lints

import (
	"testing"
)

func TestSubCertKeyUsageWithoutBits(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageNoBits.pem"
	expected := Error
	out := Lints["e_ext_key_usage_without_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSubCertKeyUsageWithBits(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	expected := Pass
	out := Lints["e_ext_key_usage_without_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSubCertKeyUsageNotIncludedBits(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.pem"
	expected := NA
	out := Lints["e_ext_key_usage_without_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
