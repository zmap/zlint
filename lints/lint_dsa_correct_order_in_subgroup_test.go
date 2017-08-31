package lints

import (
	"crypto/dsa"
	"math/big"
	"testing"
)

func TestDSACorrectOrderSubgroup(t *testing.T) {
	inputPath := "../testlint/testCerts/dsaCorrectOrderInSubgroup.pem"
	expected := Pass
	out := Lints["e_dsa_correct_order_in_subgroup"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDSANotCorrectOrderSubgroup(t *testing.T) {
	inputPath := "../testlint/testCerts/dsaCorrectOrderInSubgroup.pem"
	c := ReadCertificate(inputPath)
	dsaKey := c.PublicKey.(*dsa.PublicKey)
	pMinusOne := big.NewInt(0)
	pMinusOne.Sub(dsaKey.P, big.NewInt(1))
	dsaKey.Y = pMinusOne
	expected := Error
	out := Lints["e_dsa_correct_order_in_subgroup"].Execute(c)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


