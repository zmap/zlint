// lint_br_ian_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrIanWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/ianWildcardFirst.cer"
	desEnum := Error
	out, _ := Lints["e_ian_wildcard_not_first"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIanWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_ian_wildcard_not_first"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
