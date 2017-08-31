package lints

import (
	"testing"
)

func TestRootCAKeyUsageCritical(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAKeyUsagePresent.pem"
	expected := Pass
	out := Lints["e_root_ca_key_usage_must_be_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRootCAKeyUsageNotCritical(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAKeyUsageNotCritical.pem"
	expected := Error
	out := Lints["e_root_ca_key_usage_must_be_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
