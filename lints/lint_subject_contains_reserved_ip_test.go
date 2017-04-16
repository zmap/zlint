// lint_subject_contains_reserved_ip_test.go
package lints

import (
	"testing"
)

func TestSubjectIPReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectReservedIP.pem"
	desEnum := Error
	out, _ := Lints["e_subject_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectIPReserved6(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectReservedIP6.pem"
	desEnum := Error
	out, _ := Lints["e_subject_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectIPNotReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectGoodIP.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
