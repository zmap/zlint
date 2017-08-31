// lint_ext_policy_map_not_in_cert_policy_test.go
package lints

import (
	"testing"
)

func TestPolicyMapInCertPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapIssuerNotInCertPolicy.pem"
	desEnum := Warn
	out := Lints["w_ext_policy_map_not_in_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPolicyMapNotInCertPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapGood.pem"
	desEnum := Pass
	out := Lints["w_ext_policy_map_not_in_cert_policy"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
