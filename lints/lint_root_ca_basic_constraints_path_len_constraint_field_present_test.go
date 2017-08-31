// lint_root_ca_basic_constraints_path_len_constraint_field_present_test.go
package lints

import (
	"testing"
)

func TestRootCaMaxLenPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCaMaxPathLenPresent.pem"
	expected := Warn
	out := Lints["w_root_ca_basic_constraints_path_len_constraint_field_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRootCaMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCaMaxPathLenMissing.pem"
	expected := Pass
	out := Lints["w_root_ca_basic_constraints_path_len_constraint_field_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
