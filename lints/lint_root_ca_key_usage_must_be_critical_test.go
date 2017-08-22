package lints

import (
	"testing"
)

func TestRootCAKeyUsageCritical(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAKeyUsagePresent.pem"
	desEnum := Pass
	out, _ := Lints["e_root_ca_key_usage_must_be_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRootCAKeyUsageNotCritical(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAKeyUsageNotCritical.pem"
	desEnum := Error
	out, _ := Lints["e_root_ca_key_usage_must_be_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
