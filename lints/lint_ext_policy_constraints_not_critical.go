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

type policyConstraintsCritical struct {
	// Internal data here
}

func (l *policyConstraintsCritical) Initialize() error {
	return nil
}

func (l *policyConstraintsCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.PolicyConstOID)
}

func (l *policyConstraintsCritical) RunTest(c *x509.Certificate) (ResultStruct, error) {
	pc := util.GetExtFromCert(c, util.PolicyConstOID)
	if !pc.Critical {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_policy_constraints_not_critical",
		Description:   "Conforming CAs MUST mark the policy constraints extension as critical",
		Source:        "RFC 5280: 4.2.1.11",
		EffectiveDate: util.RFC5280Date,
		Test:          &policyConstraintsCritical{},
	})
}
