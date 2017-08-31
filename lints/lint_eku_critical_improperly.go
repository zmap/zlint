// lint_eku_critical_improperly.go
/************************************************
RFC 5280: 4.2.1.12
If a CA includes extended key usages to satisfy such applications,
   but does not wish to restrict usages of the key, the CA can include
   the special KeyPurposeId anyExtendedKeyUsage in addition to the
   particular key purposes required by the applications.  Conforming CAs
   SHOULD NOT mark this extension as critical if the anyExtendedKeyUsage
   KeyPurposeId is present.  Applications that require the presence of a
   particular purpose MAY reject certificates that include the
   anyExtendedKeyUsage OID but not the particular OID expected for the
   application.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ekuBadCritical struct {
	// Internal data here
}

func (l *ekuBadCritical) Initialize() error {
	return nil
}

func (l *ekuBadCritical) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsExtInCert(c, util.EkuSynOid)
}

func (l *ekuBadCritical) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if e := util.GetExtFromCert(c, util.EkuSynOid); e.Critical {
		for _, single_use := range c.ExtKeyUsage {
			if single_use == x509.ExtKeyUsageAny {
				return ResultStruct{Result: Warn}, nil
			}
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_eku_critical_improperly",
		Description:   "Conforming CAs SHOULD NOT mark extended key usage extension as critical if the anyExtendedKeyUsage KeyPurposedID is present",
		Source:        "RFC 5280: 4.2.1.12",
		EffectiveDate: util.RFC3280Date,
		Test:          &ekuBadCritical{},
	})
}
