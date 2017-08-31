// lint_ext_crl_distribution_marked_critical.go
/************************************************
The CRL distribution points extension identifies how CRL information is obtained. The extension SHOULD be non-critical, but this profile RECOMMENDS support for this extension by CAs and applications.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtCrlDistributionMarkedCritical struct{}

func (l *ExtCrlDistributionMarkedCritical) Initialize() error {
	return nil
}

func (l *ExtCrlDistributionMarkedCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.CrlDistOID)
}

func (l *ExtCrlDistributionMarkedCritical) Execute(cert *x509.Certificate) LintResult {
	if e := util.GetExtFromCert(cert, util.CrlDistOID); e != nil {
		if e.Critical == false {
			return &LintResult{Status: Pass}
		} else {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: NA}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_crl_distribution_marked_critical",
		Description:   "If included, the CRL Distribution Points extension SHOULD NOT be marked critical",
		Source:        "RFC 5280: 4.2.1.13",
		EffectiveDate: util.RFC2459Date,
		Lint:          &ExtCrlDistributionMarkedCritical{},
	})
}
