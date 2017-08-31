// lint_ext_cert_policy_explicit_text_ia5_string_test.go
package lints

import (
	"testing"
)

func TestExplicitTextIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticePres.pem"
	desEnum := Error
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextNotIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeExpTextNotIA5String.pem"
	desEnum := Pass
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextNotPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	desEnum := NA
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextNotPresent2(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeUnrecommended.pem"
	desEnum := NA
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
