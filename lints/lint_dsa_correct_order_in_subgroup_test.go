package lints

import (
	"crypto/dsa"
	"math/big"
	"testing"
)

func TestDSACorrectOrderSubgroup(t *testing.T) {

	inputPath := "../testlint/testCerts/dsaCorrectOrderInSubgroup.pem"
	desEnum := Pass
	out := Lints["e_dsa_correct_order_in_subgroup"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestDSANotCorrectOrderSubgroup(t *testing.T) {
	inputPath := "../testlint/testCerts/dsaCorrectOrderInSubgroup.pem"
	c := ReadCertificate(inputPath)
	dsaKey := c.PublicKey.(*dsa.PublicKey)
	pMinusOne := big.NewInt(0)
	pMinusOne.Sub(dsaKey.P, big.NewInt(1))
	dsaKey.Y = pMinusOne
	desEnum := Error
	out := Lints["e_dsa_correct_order_in_subgroup"].Execute(c)
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
