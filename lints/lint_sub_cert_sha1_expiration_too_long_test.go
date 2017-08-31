// lint_sub_cert_sha1_expiration_too_long_test.go
package lints

import (
	"testing"
)

func TestRsaSha1TooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/sha1ExpireAfter2017.pem"
	desEnum := Warn
	out := Lints["w_sub_cert_sha1_expiration_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRsaSha1NotTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/sha1ExpirePrior2017.pem"
	desEnum := Pass
	out := Lints["w_sub_cert_sha1_expiration_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
