// lint_ext_san_no_entries_test.go
package lints

import (
	"testing"
)

func TestSANNoEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/SANNoEntries.pem"
	desEnum := Error
	out, _ := Lints["e_ext_san_no_entries"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANHasEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_san_no_entries"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
