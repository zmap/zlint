package lints

import (
	"crypto/dsa"
	"math/big"
	"testing"
)

func TestDSACorrectOrderSubgroup(t *testing.T) {

	inputPath := "../testlint/testCerts/dsaCorrectOrderInSubgroup.pem"
	desEnum := Pass
	out, _ := Lints["e_dsa_correct_order_in_subgroup"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
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
	out, _ := Lints["e_dsa_correct_order_in_subgroup"].ExecuteTest(c)
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
