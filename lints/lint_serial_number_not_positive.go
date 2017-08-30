// lint_serial_number_not_positive.go
/************************************************
4.1.2.2.  Serial Number
   The serial number MUST be a positive integer assigned by the CA to each
   certificate. It MUST be unique for each certificate issued by a given CA
   (i.e., the issuer name and serial number identify a unique certificate).
   CAs MUST force the serialNumber to be a non-negative integer.

   Given the uniqueness requirements above, serial numbers can be expected to
   contain long integers.  Certificate users MUST be able to handle serialNumber
   values up to 20 octets.  Conforming CAs MUST NOT use serialNumber values longer
   than 20 octets.

   Note: Non-conforming CAs may issue certificates with serial numbers that are
   negative or zero.  Certificate users SHOULD be prepared togracefully handle
   such certificates.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SerialNumberNotPositive struct {
	// Internal data here
}

func (l *SerialNumberNotPositive) Initialize() error {
	return nil
}

func (l *SerialNumberNotPositive) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *SerialNumberNotPositive) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if cert.SerialNumber.Sign() == -1 { // -1 Means negative when using big.Sign()
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_serial_number_not_positive",
		Description:   "Certificates must have a positive serial number",
		Source:        "RFC 5280: 4.1.2.2",
		EffectiveDate: util.RFC3280Date,
		Test:          &SerialNumberNotPositive{},
	})
}
