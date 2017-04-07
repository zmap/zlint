// lint_signature_algorithm_weak_test.go

package lints

import (
	"testing"
)

func TestRSASHA1Weak(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Weak.cer"
	desEnum := Warn
	out, _ := Lints["w_signature_algorithm_weak"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRSASHA1Good(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.cer"
	desEnum := Pass
	out, _ := Lints["w_signature_algorithm_weak"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRSASHA256Good(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA256Good.cer"
	desEnum := Pass
	out, _ := Lints["w_signature_algorithm_weak"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRSASHA256Weak(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA256Weak.cer"
	desEnum := Warn
	out, _ := Lints["w_signature_algorithm_weak"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}


