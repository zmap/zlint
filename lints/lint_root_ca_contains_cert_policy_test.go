// lint_root_ca_contains_cert_policy_test.go
package lints

import (
	"testing"
)

func TestRootCACertPolicy(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAWithEKUCertPolicy.pem"
	desEnum := Warn
	out, _ := Lints["w_root_ca_contains_cert_policy"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRootCANoCertPolicy(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCAValid.pem"
	desEnum := Pass
	out, _ := Lints["w_root_ca_contains_cert_policy"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
