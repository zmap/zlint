// lint_distribution_point_missing_ldap_or_uri_test.go
package lints

import (
	"testing"
)

func TestDNSNameWildcardOnlyInLeftLabel(t *testing.T) {
	
	inputPath := "../testlint/testCerts/dnsNameWildcardOnlyInLeftLabel.pem"
	expected := Pass
	out := Lints["e_dnsname_wildcard_only_in_left_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestDNSNameWildcardNotOnlyInLeftLabel(t *testing.T) {
	
	inputPath := "../testlint/testCerts/dnsNameWildcardNotOnlyInLeftLabel.pem"
	expected := Error
	out := Lints["e_dnsname_wildcard_only_in_left_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
