// lint_ext_san_contains_reserved_ip_test.go
package lints

import (

  "testing"
)

func TestSanIPReserved(t *testing.T) {
  inputPath := "../testlint/testCerts/sanReservedIP.cer"
  desEnum := Error
  out, _ := Lints["ext_san_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
  if out.Result != desEnum {
    t.Error(
            "For", inputPath, /* input path*/
            "expected", desEnum,  /* The enum you expected */
            "got", out.Result, /* Actual Result */
          )
  }
}

func TestSanIPReserved6(t *testing.T) {
  inputPath := "../testlint/testCerts/sanReservedIP6.cer"
  desEnum := Error
  out, _ := Lints["ext_san_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
  if out.Result != desEnum {
    t.Error(
            "For", inputPath, /* input path*/
            "expected", desEnum,  /* The enum you expected */
            "got", out.Result, /* Actual Result */
          )
  }
}

func TestSanIPNotReserved(t *testing.T) {
  inputPath := "../testlint/testCerts/sanValidIP.cer"
  desEnum := Pass
  out, _ := Lints["ext_san_contains_reserved_ip"].ExecuteTest(ReadCertificate(inputPath))
  if out.Result != desEnum {
    t.Error(
            "For", inputPath, /* input path*/
            "expected", desEnum,  /* The enum you expected */
            "got", out.Result, /* Actual Result */
          )
  }
}
