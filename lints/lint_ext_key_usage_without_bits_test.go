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
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageWithBits(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	expected := Pass
	out := Lints["e_ext_key_usage_without_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageNotIncludedBits(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.pem"
	expected := NA
	out := Lints["e_ext_key_usage_without_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
