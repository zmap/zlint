// lint_ext_key_usage_without_bits_test.go
package lints

import (
	"testing"
)

func TestSubCertKeyUsageWithoutBits(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageNoBits.cer"
	desEnum := Error
	out, _ := Lints["e_ext_key_usage_without_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageWithBits(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_key_usage_without_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageNotIncludedBits(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.cer"
	desEnum := NA
	out, _ := Lints["e_ext_key_usage_without_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
