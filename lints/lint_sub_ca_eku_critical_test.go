// lint_sub_ca_eku_critical_test.go
package lints

import (

	"testing"
)

func TestSubCaEkuCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuCrit.cer"
	desEnum := Warn
	out, _ := Lints["sub_ca_eku_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaEkuNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuNoCrit.cer"
	desEnum := Pass
	out, _ := Lints["sub_ca_eku_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
