// lint_issuer_multiple_attr_in_rdn_test.go

package lints

import (
	"testing"
)

func TestIssuerRDNTwoAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerRDNTwoAttribute.pem"
	desEnum := Warn
	out, _ := Lints["w_issuer_multiple_attr_in_rdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIssuerRDNOneAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.pem"
	desEnum := Pass
	out, _ := Lints["w_issuer_multiple_attr_in_rdn"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
