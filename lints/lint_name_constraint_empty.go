// lint_name_constraint_empty.go
/***********************************************************************
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
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type nameConstraintEmpty struct {
	// Internal data here
}

func (l *nameConstraintEmpty) Initialize() error {
	return nil
}

func (l *nameConstraintEmpty) CheckApplies(c *x509.Certificate) bool {
	if !(util.IsExtInCert(c, util.NameConstOID)) {
		return false
	}
	nc := util.GetExtFromCert(c, util.NameConstOID)
	var seq asn1.RawValue
	rest, err := asn1.Unmarshal(nc.Value, &seq) //only one sequence, so rest should be empty
	if err != nil || len(rest) != 0 || seq.Tag != 16 || seq.Class != 0 || !seq.IsCompound {
		return false
	}
	return true
}

func (l *nameConstraintEmpty) RunTest(c *x509.Certificate) (ResultStruct, error) {
	nc := util.GetExtFromCert(c, util.NameConstOID)
	var seq asn1.RawValue
	_, err := asn1.Unmarshal(nc.Value, &seq) //only one sequence, so rest should be empty
	if err != nil {
		return ResultStruct{Result: Fatal}, nil
	}
	if len(seq.Bytes) == 0 {
		return ResultStruct{Result: Error}, nil
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_name_constraint_empty",
		Description:   "Conforming CAs MUST NOT issue certificates where name constraints is an empty sequence. That is, either the permittedSubtree or excludedSubtree fields must be present",
		Source:        "RFC 5280: 4.2.1.10",
		EffectiveDate: util.RFC5280Date,
		Test:          &nameConstraintEmpty{},
	})
}
