// lint_root_ca_extended_key_usage_present_test.go
package lints

import (
	"testing"
)

func TestRootCAEKU(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAWithEKU.pem"
	desEnum := Error
	out := Lints["e_root_ca_extended_key_usage_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRootCANoEKU(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAValid.pem"
	desEnum := Pass
	out := Lints["e_root_ca_extended_key_usage_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
