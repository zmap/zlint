// lint_ext_aia_access_location_missing_test.go
package lints

import (
	"testing"
)

func TestAIAcaIssuerMissingHTTPorLDAP(t *testing.T) {
	
	inputPath := "../testlint/testCerts/caIssuerNoHTTPLDAP.pem"
	expected := Warn
	out := Lints["w_ext_aia_access_location_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestAIAcaIssuerHTTP(t *testing.T) {
	
	inputPath := "../testlint/testCerts/caIssuerHTTP.pem"
	expected := Pass
	out := Lints["w_ext_aia_access_location_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestAIAcaIssuerLDAP(t *testing.T) {
	
	inputPath := "../testlint/testCerts/caIssuerLDAP.pem"
	expected := Pass
	out := Lints["w_ext_aia_access_location_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestAIAcaIssuerBlank(t *testing.T) {
	
	inputPath := "../testlint/testCerts/caIssuerBlank.pem"
	expected := NA
	out := Lints["w_ext_aia_access_location_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
