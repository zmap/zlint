// lint_ca_country_name_invalid_test.go
package lints

import (
	"testing"
)

func TestCaCommonNameMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caCommonNameMissing.pem"
	expected := Error
	out := Lints["e_ca_common_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCaCommonNameNotMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caCommonNameNotMissing.pem"
	expected := Pass
	out := Lints["e_ca_common_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
