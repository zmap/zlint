// lint_br_SAN_bare_wildcard_test.go
package lints

import (
	"testing"
)

func TestBrSanBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANBareWildcard.cer"
	desEnum := Error
	out, _ := Lints["e_SAN_bare_wildcard"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrSanNotBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_SAN_bare_wildcard"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
