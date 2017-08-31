// lint_ext_ian_no_entries_test.go
package lints

import (
	"testing"
)

func TestIANNoEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmpty.pem"
	expected := Error
	out := Lints["e_ext_ian_no_entries"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANHasEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSIA5String.pem"
	expected := Pass
	out := Lints["e_ext_ian_no_entries"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
