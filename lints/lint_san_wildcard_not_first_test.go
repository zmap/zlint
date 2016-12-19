// lint_br_san_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrSanWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/sanWildcardFirst.cer"
	desEnum := Error
	out, _ := Lints["san_wildcard_not_first"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrSanWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURIValid.cer"
	desEnum := Pass
	out, _ := Lints["san_wildcard_not_first"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
