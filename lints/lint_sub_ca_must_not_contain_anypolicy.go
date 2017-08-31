package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCaMustNotContainAnyPolicy struct {
	// Internal data here
}

func (l *subCaMustNotContainAnyPolicy) Initialize() error {
	return nil
}

func (l *subCaMustNotContainAnyPolicy) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c)
}

func (l *subCaMustNotContainAnyPolicy) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, policy := range c.PolicyIdentifiers {
		if policy.Equal(util.AnyPolicyOID) {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_must_not_contain_any_policy",
		Description:   "Subordinate CA: MUST NOT contain the anyPolicy identifier (2.5.29.32.0)",
		Source:        "BRs: 7.1.6.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCaMustNotContainAnyPolicy{},
	})
}
