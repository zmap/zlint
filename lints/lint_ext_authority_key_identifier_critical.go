// lint_ext_authority_key_identifier_critical.go
/*********************************************************
RFC 5280: 4.2.1.1
Conforming CAs MUST mark this extension as non-critical.
**********************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type authorityKeyIdCritical struct {
	// Internal data here
}

func (l *authorityKeyIdCritical) Initialize() error {
	return nil
}

func (l *authorityKeyIdCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.AuthkeyOID)
}

func (l *authorityKeyIdCritical) RunTest(c *x509.Certificate) (ResultStruct, error) {
	aki := util.GetExtFromCert(c, util.AuthkeyOID) //pointer to the extension
	if aki.Critical {
		return ResultStruct{Result: Error}, nil
	} else { //implies !aki.Critical
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_authority_key_identifier_critical",
		Description:   "The authority key identifier extension must be non-critical",
		Source:        "RFC 5280: 4.2.1.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &authorityKeyIdCritical{},
	})
}
