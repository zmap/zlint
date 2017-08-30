// lint_sub_key_usage_crl_sign_bit_set.go
/**************************************************************************
BRs: 7.1.2.3
keyUsage (optional)
If present, bit positions for keyCertSign and cRLSign MUST NOT be set.
***************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCrlSignAllowed struct {
	// Internal data here
}

func (l *subCrlSignAllowed) Initialize() error {
	return nil
}

func (l *subCrlSignAllowed) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsExtInCert(c, util.KeyUsageOID) && !util.IsCACert(c)
}

func (l *subCrlSignAllowed) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	if (c.KeyUsage & x509.KeyUsageCRLSign) == x509.KeyUsageCRLSign {
		return ResultStruct{Result: Error}, nil
	} else { //key usage doesn't allow cert signing or isn't present
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_key_usage_crl_sign_bit_set",
		Description:   "Subscriber Certificate: keyUsage if present, bit positions for keyCertSign and cRLSign MUST NOT be set.",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCrlSignAllowed{},
	})
}
