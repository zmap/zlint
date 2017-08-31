// lint_subject_common_name_included_test.go
package lints

import (
	"testing"
)

func TestCN(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNamesURL.pem"
	desEnum := Notice
	out := Lints["n_subject_common_name_included"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNoCN(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNamesGood.pem"
	desEnum := Pass
	out := Lints["n_subject_common_name_included"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
