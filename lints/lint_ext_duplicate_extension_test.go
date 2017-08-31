// lint_ext_duplicate_extension_test.go
package lints

import (
	"testing"
)

func TestDuplicateExtension(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/extSANDuplicated.pem"
	expected := Error
	out := Lints["e_ext_duplicate_extension"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNoDuplicateExtension(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caBasicConstCrit.pem"
	expected := Pass
	out := Lints["e_ext_duplicate_extension"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
