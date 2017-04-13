// lint_ext_key_usage_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCertKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageNotCriticalSubCert.pem"
	desEnum := Warn
	out, _ := Lints["w_ext_key_usage_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNotCrit.pem"
	desEnum := Warn
	out, _ := Lints["w_ext_key_usage_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out, _ := Lints["w_ext_key_usage_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	desEnum := Pass
	out, _ := Lints["w_ext_key_usage_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageNotIncludedCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.pem"
	desEnum := NA
	out, _ := Lints["e_ext_key_usage_without_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
