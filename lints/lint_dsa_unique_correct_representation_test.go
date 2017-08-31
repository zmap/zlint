package lints

import (
	"crypto/dsa"
	"math/big"
	"testing"
)

func TestDSAUniqueCorrectRepresentation(t *testing.T) {
	inputPath := "../testlint/testCerts/dsaUniqueRep.pem"
	desEnum := Pass
	out := Lints["e_dsa_unique_correct_representation"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath,
			"expected", desEnum,
			"got", out.Status,
		)
	}
}

func TestDSANotUniqueCorrectRepresentation(t *testing.T) {
	inputPath := "../testlint/testCerts/dsaUniqueRep.pem"
	c := ReadCertificate(inputPath)

	// Replace Y with P - 1
	dsaKey := c.PublicKey.(*dsa.PublicKey)
	pMinusOne := big.NewInt(0)
	pMinusOne.Sub(dsaKey.P, big.NewInt(1))
	dsaKey.Y = pMinusOne

	// Expect failure
	expected := Error
	out := Lints["e_dsa_unique_correct_representation"].Execute(c)
	if out.Status != expected {
		t.Error(
			"For", inputPath,
			"expected", expected,
			"got", out.Status,
		)
	}
}
