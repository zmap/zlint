// lint_ext_san_no_entries_test.go
package lints

import (
	"testing"
)

func TestSANNoEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/SANNoEntries.pem"
	expected := Error
	out := Lints["e_ext_san_no_entries"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANHasEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_ext_san_no_entries"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
