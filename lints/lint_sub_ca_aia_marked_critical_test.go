package lints

import (
	"testing"
)

func TestSubCAAIAMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAMarkedCritical.pem"
	expected := Error
	out := Lints["e_sub_ca_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCAAIANotMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIANotMarkedCritical.pem"
	expected := Pass
	out := Lints["e_sub_ca_aia_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
