// lint_ext_policy_constraints_not_critical.go
/************************************************
RFC 5280: 4.2.1.11
Conforming CAs MUST mark this extension as critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type policyConstraintsCritical struct{}

func (l *policyConstraintsCritical) Initialize() error {
	return nil
}

func (l *policyConstraintsCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.PolicyConstOID)
}

func (l *policyConstraintsCritical) Execute(c *x509.Certificate) *LintResult {
	pc := util.GetExtFromCert(c, util.PolicyConstOID)
	if !pc.Critical {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_policy_constraints_not_critical",
		Description:   "Conforming CAs MUST mark the policy constraints extension as critical",
		Citation:      "RFC 5280: 4.2.1.11",
		Source:        RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &policyConstraintsCritical{},
	})
}
