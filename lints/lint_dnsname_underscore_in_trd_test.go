// lint_distribution_point_missing_ldap_or_uri_test.go
package lints

import (
	"testing"
)

func TestDNSNameUnderscoreInTRD(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameUnderscoreInTRD.pem"
	desEnum := Warn
	out := Lints["w_dnsname_underscore_in_trd"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDNSNameNoUnderscoreInTRD(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameNoUnderscoreInTRD.pem"
	desEnum := Pass
	out := Lints["w_dnsname_underscore_in_trd"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
