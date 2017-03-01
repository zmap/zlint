// lint_ext_duplicate_extension_test.go
package lints

import (
	"testing"
)

func TestDuplicateExtension(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/extSANDuplicated.cer"
	desEnum := Error
	out, _ := Lints["ext_duplicate_extension"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNoDuplicateExtension(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caBasicConstCrit.cer"
	desEnum := Pass
	out, _ := Lints["ext_duplicate_extension"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
