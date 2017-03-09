// lint_ext_san_registered_id_present_test.go
package lints

import (
	"testing"
)

func TestSanRegIdMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/sanCaGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_registered_id_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanRegIdPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/sanRegisteredIdBeginning.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_registered_id_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanRegIdPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/sanRegisteredIdEnd.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_registered_id_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
