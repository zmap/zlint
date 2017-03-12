// lint_ext_aia_access_location_missing_test.go
package lints

import (
	"testing"
)

func TestAIAcaIssuerMissingHTTPorLDAP(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caIssuerNoHTTPLDAP.cer"
	desEnum := Warn
	out, _ := Lints["w_ext_aia_access_location_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestAIAcaIssuerHTTP(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caIssuerHTTP.cer"
	desEnum := Pass
	out, _ := Lints["w_ext_aia_access_location_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestAIAcaIssuerLDAP(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caIssuerLDAP.cer"
	desEnum := Pass
	out, _ := Lints["w_ext_aia_access_location_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestAIAcaIssuerBlank(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caIssuerBlank.cer"
	desEnum := NA
	out, _ := Lints["w_ext_aia_access_location_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
