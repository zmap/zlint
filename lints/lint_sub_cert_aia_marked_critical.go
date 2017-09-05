package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertAiaMarkedCritical struct{}

func (l *subCertAiaMarkedCritical) Initialize() error {
	return nil
}

func (l *subCertAiaMarkedCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.IsExtInCert(c, util.AiaOID)
}

func (l *subCertAiaMarkedCritical) Execute(c *x509.Certificate) *LintResult {
	e := util.GetExtFromCert(c, util.AiaOID)
	if e.Critical {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_aia_marked_critical",
		Description:   "Subscriber Certificate: authorityInformationAccess MUST NOT be marked critical",
		Source:        "BRs: 7.1.2.3",
		Type:          BRs,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCertAiaMarkedCritical{},
	})
}
