package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCAEKUValidFields struct{}

func (l *subCAEKUValidFields) Initialize() error {
	return nil
}

func (l *subCAEKUValidFields) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.EkuSynOid)
}

func (l *subCAEKUValidFields) Execute(c *x509.Certificate) *LintResult {
	validFieldsPresent := false
	for _, ekuValue := range c.ExtKeyUsage {
		if ekuValue == x509.ExtKeyUsageServerAuth ||
			ekuValue == x509.ExtKeyUsageClientAuth {
			validFieldsPresent = true
		}
	}
	if validFieldsPresent {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Notice}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_sub_ca_eku_not_technically_constrained",
		Description:   "Subordinate CA extkeyUsage, either id-kp-serverAuth or id-kp-clientAuth or both values MUST be present to be technically constrained.",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.CABV116Date,
		Lint:          &subCAEKUValidFields{},
	})
}
