// lint_ext_key_usage_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCertKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageNotCriticalSubCert.cer"
	desEnum := Warn
	out, _ := Lints["ext_key_usage_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNotCrit.cer"
	desEnum := Warn
	out, _ := Lints["ext_key_usage_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.cer"
	desEnum := Pass
	out, _ := Lints["ext_key_usage_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.cer"
	desEnum := Pass
	out, _ := Lints["ext_key_usage_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageNotIncludedCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.cer"
	desEnum := NA
	out, _ := Lints["ext_key_usage_without_bits"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
