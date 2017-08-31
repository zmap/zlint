// lint_subject_common_name_disallowed_test.go
package lints

import (
	"testing"
)

func TestStreetAddressShouldNotExist(t *testing.T) {
	inputPath := "../testlint/testCerts/streetAddressCannotExist.pem"
	desEnum := Error
	out := Lints["e_sub_cert_street_address_should_not_exist"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestStreetAddressCanExist(t *testing.T) {
	inputPath := "../testlint/testCerts/streetAddressCanExist.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_street_address_should_not_exist"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
