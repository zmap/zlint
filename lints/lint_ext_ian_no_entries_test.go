// lint_ext_ian_no_entries_test.go
package lints

import (
	"testing"
)

func TestIANNoEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmpty.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_no_entries"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANHasEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSIA5String.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_no_entries"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
