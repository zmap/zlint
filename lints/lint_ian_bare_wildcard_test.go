// lint_br_ian_bare_wildcard_test.go
package lints

import (
	"testing"
)

func TestBrIanBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/ianBareWildcard.cer"
	desEnum := Error
	out, _ := Lints["ian_bare_wildcard"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIanNotBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/ianURIValid.cer"
	desEnum := Pass
	out, _ := Lints["ian_bare_wildcard"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
