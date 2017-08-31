// lint_ext_key_usage_without_bits.go
/***********************************************************************
 This profile does not restrict the combinations of bits that may be
   set in an instantiation of the keyUsage extension.  However,
   appropriate values for keyUsage extensions for particular algorithms
   are specified in [RFC3279], [RFC4055], and [RFC4491].  When the
   keyUsage extension appears in a certificate, at least one of the bits
   MUST be set to 1.
***********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type keyUsageBitsSet struct {
	// Internal data here
}

func (l *keyUsageBitsSet) Initialize() error {
	return nil
}

func (l *keyUsageBitsSet) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *keyUsageBitsSet) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.KeyUsage == 0 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_key_usage_without_bits",
		Description:   "When the keyUsage extension is included, at least one bit MUST be set to 1",
		Source:        "RFC 5280: 4.2.1.3",
		EffectiveDate: util.RFC5280Date,
		Test:          &keyUsageBitsSet{},
	})
}
