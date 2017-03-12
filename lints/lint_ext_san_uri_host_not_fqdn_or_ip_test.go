// lint_ext_san_uri_host_not_fqdn_or_ip_test.go
package lints

import (
	"testing"
)

func TestSanUriNotFqdn(t *testing.T) {
	inputPath := "../testlint/testCerts/sanUriNotFqdn.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

// func TestSanUriFqdn(t *testing.T) {
// 	inputPath := "../testlint/testCerts/sanUriFqdn.cer"
// 	desEnum := Pass
// 	out, _ := Lints["ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
// 	if out.Result != desEnum {
// 		t.Error(
// 			"For", inputPath, /* input path*/
// 			"expected", desEnum, /* The enum you expected */
// 			"got", out.Result, /* Actual Result */
// 		)
// 	}
// }

// func TestSanUriIp(t *testing.T) {
// 	inputPath := "../testlint/testCerts/sanUriIp.cer"
// 	desEnum := Pass
// 	out, _ := Lints["ext_san_uri_host_not_fqdn_or_ip"].ExecuteTest(ReadCertificate(inputPath))
// 	if out.Result != desEnum {
// 		t.Error(
// 			"For", inputPath, /* input path*/
// 			"expected", desEnum, /* The enum you expected */
// 			"got", out.Result, /* Actual Result */
// 		)
// 	}
// }
