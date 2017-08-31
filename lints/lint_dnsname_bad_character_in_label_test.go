package lints

import (
	"testing"
)

func TestBadCharacterInDNSLabel(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameBadCharacterInLabel.pem"
	expected := Error
	out := Lints["e_dnsname_bad_character_in_label"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
