// lint_sub_ca_no_dns_name_contstraints.go
/**************************************************************************************************************************
If the Subordinate CA is not allowed to issue certificates with dNSNames, then the Subordinate CA Certificate
MUST include a zero‐length dNSName in excludedSubtrees. Otherwise, the Subordinate CA Certificate MUST
include at least one dNSName in permittedSubtrees.
**************************************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCaBadDNSConstraint struct {
	// Internal data here
}

func (l *subCaBadDNSConstraint) Initialize() error {
	return nil
}

func (l *subCaBadDNSConstraint) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.NameConstOID)
}

func (l *subCaBadDNSConstraint) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if len(c.PermittedDNSNames) == 0 {
		for _, excluded := range c.ExcludedDNSNames {
			if len(excluded.Data) == 0 {
				return ResultStruct{Result: Pass}, nil
			}
		}
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_no_dns_name_constraints",
		Description:   "Subordanate CA certs must include in the name contraints extension either premitted dns names or prohibit the empty DNS name.",
		Providence:    "CAB: 7.1.5",
		EffectiveDate: util.CABV116Date,
		Test:          &subCaBadDNSConstraint{}})
}
