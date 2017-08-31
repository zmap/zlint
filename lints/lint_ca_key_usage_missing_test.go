// lint_ca_key_usage_missing_test.go
package lints

import (
	"testing"
)

func TestCaKeyUsageMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caKeyUsageMissing.pem"
	desEnum := Error
	out := Lints["e_ca_key_usage_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestKeyUsagePresent(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	desEnum := Pass
	out := Lints["e_ca_key_usage_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
