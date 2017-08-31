// lint_root_ca_contains_cert_policy_test.go
package lints

import (
	"testing"
)

func TestRootCACertPolicy(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAWithCertPolicy.pem"
	desEnum := Warn
	out := Lints["w_root_ca_contains_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRootCANoCertPolicy(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAValid.pem"
	desEnum := Pass
	out := Lints["w_root_ca_contains_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
