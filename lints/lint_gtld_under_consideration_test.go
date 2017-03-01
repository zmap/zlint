// lint_gtld_under_consideration_test.go
package lints

import (
	"testing"
)

func TestGTLDCommonNameIP(t *testing.T) {
	inputPath := "../testlint/testCerts/gtldcnip.cer"
	desEnum := Pass
	out, _ := Lints["gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestGTLDCommonNameNotDN(t *testing.T) {
	inputPath := "../testlint/testCerts/gtldcnnotdn.cer"
	desEnum := Pass
	out, _ := Lints["gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestGTLDCommonNameValid(t *testing.T) {
	inputPath := "../testlint/testCerts/gtldcnvalid.cer"
	desEnum := Pass
	out, _ := Lints["gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestGTLDCommonNameInvalid(t *testing.T) {
	inputPath := "../testlint/testCerts/gtldcnbad.cer"
	desEnum := Warn
	out, _ := Lints["gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGTLDDNSNamesIP(t *testing.T) {
	inputPath := "../testlint/testCerts/gtlddnsip.cer"
	desEnum := Pass
	out, _ := Lints["gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestGTLDDNSNamesValid(t *testing.T) {
	inputPath := "../testlint/testCerts/gtlddnsvalid.cer"
	desEnum := Pass
	out, _ := Lints["gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGTLDDNSNamesNotDn(t *testing.T) {
	inputPath := "../testlint/testCerts/gtlddnsnotdn.cer"
	desEnum := Pass
	out, _ := Lints["gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGTLDDNSNamesInvalid(t *testing.T) {
	inputPath := "../testlint/testCerts/gtlddnsbad.cer"
	desEnum := Warn
	out, _ := Lints["gtld_under_consideration"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
