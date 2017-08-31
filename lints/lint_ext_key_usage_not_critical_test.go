// lint_ext_key_usage_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCertKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageNotCriticalSubCert.pem"
	desEnum := Warn
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNotCrit.pem"
	desEnum := Warn
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCaKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	desEnum := Pass
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertKeyUsageNotIncludedCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.pem"
	desEnum := NA
	out := Lints["e_ext_key_usage_without_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
