// lint_serial_number_longer_than_20_octets.go
/************************************************
RFC 5280: 4.1.2.2.  Serial Number
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

type serialNumberTooLong struct {
	// Internal data here
}

func (l *serialNumberTooLong) Initialize() error {
	return nil
}

func (l *serialNumberTooLong) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *serialNumberTooLong) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.SerialNumber.BitLen() > 160 { // 20 octets
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_serial_number_longer_than_20_octets",
		Description:   "Certificates must not have a serial number longer than 20 octets",
		Source:        "RFC 5280: 4.1.2.2",
		EffectiveDate: util.RFC3280Date,
		Test:          &serialNumberTooLong{},
	})
}
