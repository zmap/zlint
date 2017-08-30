// lint_ext_policy_map_not_in_cert_policy.go
/*********************************************************************
RFC 5280: 4.2.1.5
Each issuerDomainPolicy named in the policy mapping extension SHOULD
   also be asserted in a certificate policies extension in the same
   certificate.  Policies SHOULD NOT be mapped either to or from the
   special value anyPolicy (section 4.2.1.5).
*********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type policyMapMatchesCertPolicy struct {
	// Internal data here
}

func (l *policyMapMatchesCertPolicy) Initialize() error {
	return nil
}

func (l *policyMapMatchesCertPolicy) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.PolicyMapOID)
}

func (l *policyMapMatchesCertPolicy) RunTest(c *x509.Certificate) (ResultStruct, error) {
	extPolMap := util.GetExtFromCert(c, util.PolicyMapOID)
	polMap, err := util.GetMappedPolicies(extPolMap)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	for _, pair := range polMap {
		if !util.SliceContainsOID(c.PolicyIdentifiers, pair[0]) {
			return ResultStruct{Result: Warn}, err
		}
	}
	//else
	return ResultStruct{Result: Pass}, err
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_policy_map_not_in_cert_policy",
		Description:   "Each issuerDomainPolicy named in the policy mappings extension should also be asserted in a certificate policies extension",
		Source:        "RFC 5280: 4.2.1.5",
		EffectiveDate: util.RFC3280Date,
		Test:          &policyMapMatchesCertPolicy{},
	})
}
