package lints

import "testing"

func TestNonEmptyPermittedDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/nonEmptyPermittedDNS.pem"
	expected := Pass
	out := Lints["e_sub_ca_eku_name_constraints"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestBadExclude(t *testing.T) {
	inputPath := "../testlint/testCerts/nameConstraintsMissing.pem"
	expected := Error
	out := Lints["e_sub_ca_eku_name_constraints"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
