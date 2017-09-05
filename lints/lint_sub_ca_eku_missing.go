package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCAEKUMissing struct{}

func (l *subCAEKUMissing) Initialize() error {
	return nil
}

func (l *subCAEKUMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c)
}

func (l *subCAEKUMissing) Execute(c *x509.Certificate) *LintResult {
	if util.IsExtInCert(c, util.EkuSynOid) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_eku_missing",
		Description:   "Subordinate CA certificate MUST have extkeyUsage extension",
		Source:        "BRs: 7.1.5",
		Type:          CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCAEKUMissing{},
	})
}
