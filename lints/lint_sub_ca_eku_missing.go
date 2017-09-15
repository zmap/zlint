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
		return &LintResult{Status: Notice}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_sub_ca_eku_missing",
		Description:   "To be considered Technically Constrained, the Subordinate CA certificate MUST have extkeyUsage extension",
		Citation:      "BRs: 7.1.5",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCAEKUMissing{},
	})
}
