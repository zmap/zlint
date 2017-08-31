// lint_ext_san_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestSANGood(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_ext_san_space_dns_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithSpaceDNS.pem"
	expected := Error
	out := Lints["e_ext_san_space_dns_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
