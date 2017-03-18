// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSanDirNamePresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDirectoryNameBeginning.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_directory_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanDirNamePresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDirectoryNameEnd.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_directory_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanDirNameMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_directory_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
