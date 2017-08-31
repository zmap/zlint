// lint_ext_policy_map_any_policy.go
/********************************************************************
RFC 5280: 4.2.1.5
Each issuerDomainPolicy named in the policy mappings extension SHOULD
   also be asserted in a certificate policies extension in the same
   certificate.  Policies MUST NOT be mapped either to or from the
   special value anyPolicy (Section 4.2.1.4).
********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type policyMapAnyPolicy struct {
	// Internal data here
}

func (l *policyMapAnyPolicy) Initialize() error {
	return nil
}

func (l *policyMapAnyPolicy) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.PolicyMapOID)
}

func (l *policyMapAnyPolicy) RunTest(c *x509.Certificate) (ResultStruct, error) {
	extPolMap := util.GetExtFromCert(c, util.PolicyMapOID)
	polMap, err := util.GetMappedPolicies(extPolMap)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}

	for _, pair := range polMap {
		if util.AnyPolicyOID.Equal(pair[0]) || util.AnyPolicyOID.Equal(pair[1]) {
			return ResultStruct{Result: Error}, err
		}
	}
	return ResultStruct{Result: Pass}, err
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_policy_map_any_policy",
		Description:   "Policies must not be mapped to or from the anyPolicy value",
		Source:        "RFC 5280: 4.2.1.5",
		EffectiveDate: util.RFC3280Date,
		Test:          &policyMapAnyPolicy{},
	})
}
