// lint_ext_san_contains_reserved_ip_test.go
package lints

import (
	"testing"
)

func TestSANIPReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/SANReservedIP.pem"
	expected := Error
	out := Lints["e_ext_san_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSANIPReserved6(t *testing.T) {
	inputPath := "../testlint/testCerts/SANReservedIP6.pem"
	expected := Error
	out := Lints["e_ext_san_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSANIPNotReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/SANValidIP.pem"
	expected := Pass
	out := Lints["e_ext_san_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
