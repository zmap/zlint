// lint_root_ca_contains_cert_policy_test.go
package lints

import (
	"testing"
)

func TestRootCACertPolicy(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAWithCertPolicy.pem"
	expected := Warn
	out := Lints["w_root_ca_contains_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRootCANoCertPolicy(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAValid.pem"
	expected := Pass
	out := Lints["w_root_ca_contains_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
