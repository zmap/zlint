// lint_ext_name_constraints_not_critical.go
/************************************************************************
Restrictions are defined in terms of permitted or excluded name
   subtrees.  Any name matching a restriction in the excludedSubtrees
   field is invalid regardless of information appearing in the
   permittedSubtrees.  Conforming CAs MUST mark this extension as
   critical and SHOULD NOT impose name constraints on the x400Address,
   ediPartyName, or registeredID name forms.  Conforming CAs MUST NOT
   issue certificates where name constraints is an empty sequence.  That
   is, either the permittedSubtrees field or the excludedSubtrees MUST
   be present.
************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type nameConstraintCrit struct {
	// Internal data here
}

func (l *nameConstraintCrit) Initialize() error {
	return nil
}

func (l *nameConstraintCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.NameConstOID)
}

func (l *nameConstraintCrit) RunTest(c *x509.Certificate) (ResultStruct, error) {
	e := util.GetExtFromCert(c, util.NameConstOID)
	if e.Critical {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_name_constraints_not_critical",
		Description:   "If it is included, conforming CAs MUST mark the name constrains extension as critical",
		Source:        "RFC 5280: 4.2.1.10",
		EffectiveDate: util.RFC2459Date,
		Test:          &nameConstraintCrit{},
	})
}
