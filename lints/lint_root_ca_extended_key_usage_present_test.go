// lint_root_ca_extended_key_usage_present_test.go
package lints

import (
	"testing"
)

func TestRootCAEKU(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAWithEKUCertPolicy.pem"
	desEnum := Error
	out, _ := Lints["e_root_ca_extended_key_usage_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRootCANoEKU(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAValid.pem"
	desEnum := Pass
	out, _ := Lints["e_root_ca_extended_key_usage_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
