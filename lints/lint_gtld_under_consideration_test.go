// lint_gtld_under_consideration_test.go
package lints

import (
	"testing"
)

func TestGTLDCommonNameIP(t *testing.T) {
	inputPath := "../testlint/testCerts/gtldcnip.pem"
	desEnum := Pass
	out, _ := Lints["w_gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestGTLDCommonNameNotDN(t *testing.T) {
	inputPath := "../testlint/testCerts/gtldcnnotdn.pem"
	desEnum := Pass
	out, _ := Lints["w_gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestGTLDCommonNameValid(t *testing.T) {
	inputPath := "../testlint/testCerts/gtldcnvalid.pem"
	desEnum := Pass
	out, _ := Lints["w_gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestGTLDCommonNameInvalid(t *testing.T) {
	inputPath := "../testlint/testCerts/gtldcnbad.pem"
	desEnum := Warn
	out, _ := Lints["w_gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGTLDDNSNamesIP(t *testing.T) {
	inputPath := "../testlint/testCerts/gtlddnsip.pem"
	desEnum := Pass
	out, _ := Lints["w_gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestGTLDDNSNamesValid(t *testing.T) {
	inputPath := "../testlint/testCerts/gtlddnsvalid.pem"
	desEnum := Pass
	out, _ := Lints["w_gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGTLDDNSNamesNotDn(t *testing.T) {
	inputPath := "../testlint/testCerts/gtlddnsnotdn.pem"
	desEnum := Pass
	out, _ := Lints["w_gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGTLDDNSNamesInvalid(t *testing.T) {
	inputPath := "../testlint/testCerts/gtlddnsbad.pem"
	desEnum := Warn
	out, _ := Lints["w_gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
