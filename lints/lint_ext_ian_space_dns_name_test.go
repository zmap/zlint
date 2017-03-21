// lint_ext_ian_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestIANEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmptyDNS.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_space_dns_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANNotEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/IANNonEmptyDNS.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_space_dns_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
