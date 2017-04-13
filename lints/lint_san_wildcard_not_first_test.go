// lint_br_san_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrSANWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWildcardFirst.pem"
	desEnum := Error
	out, _ := Lints["e_san_wildcard_not_first"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrSANWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	desEnum := Pass
	out, _ := Lints["e_san_wildcard_not_first"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
