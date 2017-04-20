// lint_subject_dn_leading_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubjectDNLeadingSpace struct {
	// Internal data here
}

func (l *SubjectDNLeadingSpace) Initialize() error {
	return nil
}

func (l *SubjectDNLeadingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectDNLeadingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	hasSpace, err := util.DNSAttributeHasSpace(c, false)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	if hasSpace&1 != 0 {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_subject_dn_leading_whitespace",
		Description:   "AttributeValue in subject RelativeDistinguishedName sequence should not have leading whitespace",
		Providence:    "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SubjectDNLeadingSpace{}})
}
