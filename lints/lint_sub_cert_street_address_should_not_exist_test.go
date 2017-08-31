// lint_subject_common_name_disallowed_test.go
package lints

import (
	"testing"
)

func TestStreetAddressShouldNotExist(t *testing.T) {
	inputPath := "../testlint/testCerts/streetAddressCannotExist.pem"
	expected := Error
	out := Lints["e_sub_cert_street_address_should_not_exist"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestStreetAddressCanExist(t *testing.T) {
	inputPath := "../testlint/testCerts/streetAddressCanExist.pem"
	expected := Pass
	out := Lints["e_sub_cert_street_address_should_not_exist"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
