package lints

import "testing"

func TestDSAShorterThan2048Bits(t *testing.T) {

	inputPath := "../testlint/testCerts/dsaShorterThan2048Bits.pem"
	desEnum := Error
	out, _ := Lints["e_dsa_shorter_than_2048_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestDSANotShorterThan2048Bits(t *testing.T) {
	inputPath := "../testlint/testCerts/dsaNotShorterThan2048Bits.pem"
	desEnum := Pass
	out, _ := Lints["e_dsa_shorter_than_2048_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
