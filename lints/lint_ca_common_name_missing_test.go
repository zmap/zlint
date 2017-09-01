// lint_ca_country_name_invalid_test.go
package lints

import (
	"testing"
)

func TestCaCommonNameMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caCommonNameMissing.pem"
	expected := Error
	out := Lints["e_ca_common_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCommonNameNotMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caCommonNameNotMissing.pem"
	expected := Pass
	out := Lints["e_ca_common_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
