// lint_distribution_point_missing_ldap_or_uri_test.go
package lints

import (
	"testing"
)

func TestDNSNameWildcardOnlyInLeftLabel(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardOnlyInLeftLabel.pem"
	desEnum := Pass
	out := Lints["e_dnsname_wildcard_only_in_left_label"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestDNSNameWildcardNotOnlyInLeftLabel(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardNotOnlyInLeftLabel.pem"
	desEnum := Error
	out := Lints["e_dnsname_wildcard_only_in_left_label"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
