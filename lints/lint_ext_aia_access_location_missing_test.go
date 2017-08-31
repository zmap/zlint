// lint_ext_aia_access_location_missing_test.go
package lints

import (
	"testing"
)

func TestAIAcaIssuerMissingHTTPorLDAP(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caIssuerNoHTTPLDAP.pem"
	desEnum := Warn
	out := Lints["w_ext_aia_access_location_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestAIAcaIssuerHTTP(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caIssuerHTTP.pem"
	desEnum := Pass
	out := Lints["w_ext_aia_access_location_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestAIAcaIssuerLDAP(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caIssuerLDAP.pem"
	desEnum := Pass
	out := Lints["w_ext_aia_access_location_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestAIAcaIssuerBlank(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caIssuerBlank.pem"
	desEnum := NA
	out := Lints["w_ext_aia_access_location_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
