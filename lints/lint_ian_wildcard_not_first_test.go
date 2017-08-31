// lint_br_ian_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrIANWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/IANWildcardFirst.pem"
	desEnum := Error
	out := Lints["e_ian_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIANWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	desEnum := Pass
	out := Lints["e_ian_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
