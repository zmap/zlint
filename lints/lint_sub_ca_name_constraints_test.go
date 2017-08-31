package lints

import "testing"

func TestNonEmptyPermittedDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/nonEmptyPermittedDNS.pem"
	desEnum := Pass
	out := Lints["e_sub_ca_eku_name_constraints"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestBadExclude(t *testing.T) {
	inputPath := "../testlint/testCerts/nameConstraintsMissing.pem"
	desEnum := Error
	out := Lints["e_sub_ca_eku_name_constraints"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
