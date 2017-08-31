// lint_ext_key_usage_cert_sign_without_ca.go
/************************************************************************
RFC 5280: 4.2.1.9
The cA boolean indicates whether the certified public key may be used
   to verify certificate signatures.  If the cA boolean is not asserted,
   then the keyCertSign bit in the key usage extension MUST NOT be
   asserted.  If the basic constraints extension is not present in a
   version 3 certificate, or the extension is present but the cA boolean
   is not asserted, then the certified public key MUST NOT be used to
   verify certificate signatures.
************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type keyUsageCertSignNoCa struct {
	// Internal data here
}

func (l *keyUsageCertSignNoCa) Initialize() error {
	return nil
}

func (l *keyUsageCertSignNoCa) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *keyUsageCertSignNoCa) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if (c.KeyUsage & x509.KeyUsageCertSign) != 0 {
		if c.BasicConstraintsValid && util.IsCACert(c) { //CA certs may assert certtificate signing usage
			return ResultStruct{Result: Pass}, nil
		} else {
			return ResultStruct{Result: Error}, nil
		}
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_key_usage_cert_sign_without_ca",
		Description:   "if the keyCertSign bit is asserted, then the cA bit in the basic constraints extension MUST also be asserted",
		Source:        "RFC 5280: 4.2.1.3 & 4.2.1.9",
		EffectiveDate: util.RFC3280Date,
		Test:          &keyUsageCertSignNoCa{},
	})
}
