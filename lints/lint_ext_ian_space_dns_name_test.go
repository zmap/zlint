// lint_ext_ian_space_dns_name_test.go
package lints

import (

	"testing"
)

func TestIanEmptyDns(t *testing.T) {
	inputPath := "../testlint/testCerts/ianEmptyDns.cer"
	desEnum := Error
	out, _ := Lints["ext_ian_space_dns_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanNotEmptyDns(t *testing.T) {
	inputPath := "../testlint/testCerts/ianNonEmptyDns.cer"
	desEnum := Pass
	out, _ := Lints["ext_ian_space_dns_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
