// lint_ext_san_uniform_resource_identifier_present_test.go
package lints

import (
	"testing"
)

func TestSANURIMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uniform_resource_identifier_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANURIPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIBeginning.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uniform_resource_identifier_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANURIPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIEnd.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uniform_resource_identifier_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
