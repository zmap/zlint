// lint_root_ca_basic_constraints_path_len_constraint_field_present_test.go
package lints

import (
	"testing"
)

func TestRootCaMaxLenPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCaMaxPathLenPresent.pem"
	desEnum := Warn
	out, _ := Lints["w_root_ca_basic_constraints_path_len_constraint_field_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRootCaMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/rootCaMaxPathLenMissing.pem"
	desEnum := Pass
	out, _ := Lints["w_root_ca_basic_constraints_path_len_constraint_field_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
