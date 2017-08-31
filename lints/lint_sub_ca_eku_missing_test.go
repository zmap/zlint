// lint_sub_ca_eku_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaEkuMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUMissing.pem"
	desEnum := Error
	out := Lints["e_sub_ca_eku_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaEkuNotMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuCrit.pem"
	desEnum := Pass
	out := Lints["e_sub_ca_eku_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
