// lint_br_san_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrSANWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWildcardFirst.pem"
	desEnum := Error
	out := Lints["e_san_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestBrSANWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	desEnum := Pass
	out := Lints["e_san_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
