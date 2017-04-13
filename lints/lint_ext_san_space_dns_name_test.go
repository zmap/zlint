// lint_ext_san_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestSANGood(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_san_space_dns_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithSpaceDNS.pem"
	desEnum := Error
	out, _ := Lints["e_ext_san_space_dns_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
