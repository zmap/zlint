package lints

import (
	"testing"
)

func TestIANABareSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameContainsBareIANASuffix.pem"
	expected := Error
	out := Lints["e_dnsname_contains_bare_iana_suffix"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
