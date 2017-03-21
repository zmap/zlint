// lint_br_IAN_bare_wildcard_test.go
package lints

import (
	"testing"
)

func TestBrIANBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareWildcard.cer"
	desEnum := Error
	out, _ := Lints["e_IAN_bare_wildcard"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIANNotBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.cer"
	desEnum := Pass
	out, _ := Lints["e_IAN_bare_wildcard"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
