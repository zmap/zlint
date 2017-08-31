// lint_distribution_point_missing_ldap_or_uri_test.go
package lints

import (
	"testing"
)

func TestCRLDistNoHttp(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlDistribNoHTTP.pem"
	desEnum := Warn
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCRLDistHttp(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlDistribWithHTTP.pem"
	desEnum := Pass
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCRLDistLdap(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlDistribWithLDAP.pem"
	desEnum := Pass
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
