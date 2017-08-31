package lints

import (
	"testing"
)

func TestRootCAKeyUsagePresent(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAKeyUsagePresent.pem"
	desEnum := Pass
	out := Lints["e_root_ca_key_usage_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRootCAKeyUsageMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAKeyUsageMissing.pem"
	desEnum := Error
	out := Lints["e_root_ca_key_usage_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
