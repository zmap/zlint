// lint_ext_ian_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestIANEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmptyDNS.pem"
	desEnum := Error
	out := Lints["e_ext_ian_space_dns_name"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANNotEmptyDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/IANNonEmptyDNS.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_space_dns_name"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
