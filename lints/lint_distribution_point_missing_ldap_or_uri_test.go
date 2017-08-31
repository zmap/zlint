// lint_distribution_point_missing_ldap_or_uri_test.go
package lints

import (
	"testing"
)

func TestCRLDistNoHttp(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlDistribNoHTTP.pem"
	expected := Warn
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCRLDistHttp(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlDistribWithHTTP.pem"
	expected := Pass
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCRLDistLdap(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlDistribWithLDAP.pem"
	expected := Pass
	out := Lints["w_distribution_point_missing_ldap_or_uri"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
