// lint_ext_freshest_crl_marked_critical.go
/************************************************
The freshest CRL extension identifies how delta CRL information is obtained. The extension MUST be marked as non-critical by conforming CAs. Further discussion of CRL management is contained in Section 5.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
)

type ExtFreshestCrlMarkedCritical struct{}

func (l *ExtFreshestCrlMarkedCritical) Initialize() error {
	return nil
}

func (l *ExtFreshestCrlMarkedCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.FreshCRLOID)
}

func (l *ExtFreshestCrlMarkedCritical) Execute(cert *x509.Certificate) *LintResult {
	var fCRL *pkix.Extension = util.GetExtFromCert(cert, util.FreshCRLOID)
	if fCRL != nil && fCRL.Critical {
		return &LintResult{Status: Error}
	} else if fCRL != nil && !fCRL.Critical {
		return &LintResult{Status: Pass}
	}
	return &LintResult{Status: NA} //shouldn't happen
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_freshest_crl_marked_critical",
		Description:   "Freshest CRL MUST be marked as non-critical by conforming CAs",
		Source:        "RFC 5280: 4.2.1.15",
		Type:          RFC5280,
		EffectiveDate: util.RFC3280Date,
		Lint:          &ExtFreshestCrlMarkedCritical{},
	})
}
