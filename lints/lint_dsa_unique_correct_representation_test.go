// lint_distribution_point_missing_ldap_or_uri_test.go
package lints

import (
	"testing"
)

func TestDSAUniqueCorrectRepresentation(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dsaUniqueRep.pem"
	desEnum := Pass
	out, _ := Lints["e_dsa_unique_correct_representation"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
