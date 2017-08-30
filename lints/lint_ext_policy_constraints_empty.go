// lint_ext_policy_constraints_empty.go
/*************************************************************************
RFC 5280: 4.2.1.11
Conforming CAs MUST NOT issue certificates where policy constraints
   is an empty sequence.  That is, either the inhibitPolicyMapping field
   or the requireExplicitPolicy field MUST be present.  The behavior of
   clients that encounter an empty policy constraints field is not
   addressed in this profile.
*************************************************************************/

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type policyConstraintsContents struct {
	// Internal data here
}

func (l *policyConstraintsContents) Initialize() error {
	return nil
}

func (l *policyConstraintsContents) CheckApplies(c *x509.Certificate) bool {
	if !(util.IsExtInCert(c, util.PolicyConstOID)) {
		return false
	}
	pc := util.GetExtFromCert(c, util.PolicyConstOID)
	var seq asn1.RawValue
	rest, err := asn1.Unmarshal(pc.Value, &seq) //only one sequence, so rest should be empty
	if err != nil || len(rest) != 0 || seq.Tag != 16 || seq.Class != 0 || !seq.IsCompound {
		return false
	}
	return true
}

func (l *policyConstraintsContents) RunTest(c *x509.Certificate) (ResultStruct, error) {
	pc := util.GetExtFromCert(c, util.PolicyConstOID)
	var seq asn1.RawValue
	_, err := asn1.Unmarshal(pc.Value, &seq) //only one sequence, so rest should be empty
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
		Name:          "e_ext_policy_constraints_empty",
		Description:   "Conforming CAs MUST NOT issue certificates where policy constraints is an empty sequence. That is, either the inhibitPolicyMapping field or the requireExplicityPolicy field MUST be present",
		Source:        "RFC 5280: 4.2.1.11",
		EffectiveDate: util.RFC2459Date,
		Test:          &policyConstraintsContents{},
	})
}
