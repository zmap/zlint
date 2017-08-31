package lints

import (
	"testing"
)

func TestSubCAAIAMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAMarkedCritical.pem"
	desEnum := Error
	out := Lints["e_sub_ca_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCAAIANotMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIANotMarkedCritical.pem"
	desEnum := Pass
	out := Lints["e_sub_ca_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
