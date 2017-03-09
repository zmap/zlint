// lint_ext_authority_key_identifier_missing_test.go
package lints

import (
	"testing"
)

func TestAKIMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/akiMissing.cer"
	desEnum := Error
	out, _ := Lints["e_ext_authority_key_identifier_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestAKIPresent(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_authority_key_identifier_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
