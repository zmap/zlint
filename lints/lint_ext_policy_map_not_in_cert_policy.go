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

type policyMapMatchesCertPolicy struct{}

func (l *policyMapMatchesCertPolicy) Initialize() error {
	return nil
}

func (l *policyMapMatchesCertPolicy) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.PolicyMapOID)
}

func (l *policyMapMatchesCertPolicy) Execute(c *x509.Certificate) *LintResult {
	extPolMap := util.GetExtFromCert(c, util.PolicyMapOID)
	polMap, err := util.GetMappedPolicies(extPolMap)
	if err != nil {
		return &LintResult{Status: Fatal}
	}
	for _, pair := range polMap {
		if !util.SliceContainsOID(c.PolicyIdentifiers, pair[0]) {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_policy_map_not_in_cert_policy",
		Description:   "Each issuerDomainPolicy named in the policy mappings extension should also be asserted in a certificate policies extension",
		Citation:      "RFC 5280: 4.2.1.5",
		Source:        RFC5280,
		EffectiveDate: util.RFC3280Date,
		Lint:          &policyMapMatchesCertPolicy{},
	})
}
