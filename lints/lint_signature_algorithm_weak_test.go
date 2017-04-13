// lint_serial_number_too_short_test.go

package lints

import (
	"testing"
)

func TestRSASHA1Weak(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Weak.pem"
	desEnum := Warn
	out, _ := Lints["w_serial_number_too_short"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRSASHA1Good(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.pem"
	desEnum := Pass
	out, _ := Lints["w_serial_number_too_short"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRSASHA256Good(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA256Good.pem"
	desEnum := Pass
	out, _ := Lints["w_serial_number_too_short"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRSASHA256Weak(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA256Weak.pem"
	desEnum := Warn
	out, _ := Lints["w_serial_number_too_short"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
