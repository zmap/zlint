// lint_name_constraint_on_registered_id.go
/*******************************************************************
RFC 5280: 4.2.1.10
Restrictions are defined in terms of permitted or excluded name
subtrees.  Any name matching a restriction in the excludedSubtrees
field is invalid regardless of information appearing in the
permittedSubtrees.  Conforming CAs MUST mark this extension as
critical and SHOULD NOT impose name constraints on the x400Address,
ediPartyName, or registeredID name forms.  Conforming CAs MUST NOT
issue certificates where name constraints is an empty sequence.  That
is, either the permittedSubtrees field or the excludedSubtrees MUST
be present.
*******************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type nameConstraintOnRegisteredId struct{}

func (l *nameConstraintOnRegisteredId) Initialize() error {
	return nil
}

func (l *nameConstraintOnRegisteredId) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.NameConstOID)
}

func (l *nameConstraintOnRegisteredId) Execute(c *x509.Certificate) *LintResult {
	if c.PermittedRegisteredIDs != nil || c.ExcludedRegisteredIDs != nil {
		return &LintResult{Status: Warn}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_name_constraint_on_registered_id",
		Description:   "The name constraints extension SHOULD NOT impose constraints on the registeredID name form",
		Source:        "RFC 5280: 4.2.1.10",
		EffectiveDate: util.RFC5280Date,
		Lint:          &nameConstraintOnRegisteredId{},
	})
}
