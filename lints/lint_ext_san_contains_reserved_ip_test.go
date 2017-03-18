// lint_ext_san_contains_reserved_ip_test.go
package lints

import (
	"testing"
)

func TestSanIPReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/SANReservedIP.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanIPReserved6(t *testing.T) {
	inputPath := "../testlint/testCerts/SANReservedIP6.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanIPNotReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/SANValidIP.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
