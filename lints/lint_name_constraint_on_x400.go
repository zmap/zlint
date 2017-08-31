// lint_name_constraint_on_x400.go
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

type nameConstraintOnX400 struct {
	// Internal data here
}

func (l *nameConstraintOnX400) Initialize() error {
	return nil
}

func (l *nameConstraintOnX400) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.NameConstOID)
}

func (l *nameConstraintOnX400) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.PermittedX400Addresses != nil || c.ExcludedX400Addresses != nil {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_name_constraint_on_x400",
		Description:   "The name constraints extension SHOULD NOT impose constraints on the x400Address name form",
		Source:        "RFC 5280: 4.2.1.10",
		EffectiveDate: util.RFC5280Date,
		Test:          &nameConstraintOnX400{},
	})
}
