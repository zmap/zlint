package lints

import (
	"testing"
)

func TestRootCAKeyUsagePresent(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAKeyUsagePresent.pem"
	desEnum := Pass
	out, _ := Lints["e_root_ca_key_usage_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRootCAKeyUsageMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAKeyUsageMissing.pem"
	desEnum := Error
	out, _ := Lints["e_root_ca_key_usage_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
