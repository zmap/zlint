package lints

import (
	"testing"
)

func TestSubCAAIAMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAMarkedCritical.pem"
	desEnum := Error
	out, _ := Lints["e_sub_ca_aia_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCAAIANotMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIANotMarkedCritical.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_ca_aia_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}