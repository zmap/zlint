// lint_ext_san_contains_reserved_ip_test.go
package lints

import (
	"testing"
)

func TestSANIPReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/SANReservedIP.pem"
	desEnum := Error
	out := Lints["e_ext_san_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANIPReserved6(t *testing.T) {
	inputPath := "../testlint/testCerts/SANReservedIP6.pem"
	desEnum := Error
	out := Lints["e_ext_san_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANIPNotReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/SANValidIP.pem"
	desEnum := Pass
	out := Lints["e_ext_san_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
