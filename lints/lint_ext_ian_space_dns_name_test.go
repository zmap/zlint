// lint_ext_ian_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestIANEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmptyDNS.pem"
	expected := Error
	out := Lints["e_ext_ian_space_dns_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANNotEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/IANNonEmptyDNS.pem"
	expected := Pass
	out := Lints["e_ext_ian_space_dns_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
