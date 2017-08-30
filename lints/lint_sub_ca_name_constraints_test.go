package lints

import "testing"

func TestNonEmptyPermittedDNS(t *testing.T) {
	inputPath := "../testlint/testCerts/nonEmptyPermittedDNS.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_ca_eku_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBadExclude(t *testing.T) {
	inputPath := "../testlint/testCerts/nameConstraintsMissing.pem"
	desEnum := Error
	out, _ := Lints["e_sub_ca_eku_name_constraints"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
