// lint_sub_ca_name_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaNcNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNameConstNoCrit.pem"
	desEnum := Warn
	out := Lints["w_sub_ca_name_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaNcCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNameConstCrit.pem"
	desEnum := Pass
	out := Lints["w_sub_ca_name_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
