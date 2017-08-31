// lint_root_ca_basic_constraints_path_len_constraint_field_present_test.go
package lints

import (
	"testing"
)

func TestRootCaMaxLenPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCaMaxPathLenPresent.pem"
	desEnum := Warn
	out := Lints["w_root_ca_basic_constraints_path_len_constraint_field_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRootCaMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCaMaxPathLenMissing.pem"
	desEnum := Pass
	out := Lints["w_root_ca_basic_constraints_path_len_constraint_field_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
