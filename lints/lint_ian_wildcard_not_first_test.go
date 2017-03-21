// lint_br_IAN_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrIANWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/IANWildcardFirst.cer"
	desEnum := Error
	out, _ := Lints["e_IAN_wildcard_not_first"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIANWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_IAN_wildcard_not_first"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
