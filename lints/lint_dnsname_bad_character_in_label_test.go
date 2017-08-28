package lints

import (
	"testing"
)

func TestBadCharacterInDNSLabel(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameBadCharacterInLabel.pem"
	desEnum := Error
	out, _ := Lints["e_dnsname_bad_character_in_label"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
