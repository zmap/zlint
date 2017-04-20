// lint_subject_dn_trailing_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubjectDNTrailingSpace struct {
	// Internal data here
}

func (l *SubjectDNTrailingSpace) Initialize() error {
	return nil
}

func (l *SubjectDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectDNTrailingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	hasSpace, err := util.DNSAttributeHasSpace(c, false)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	if hasSpace&2 != 0 {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_subject_dn_trailing_whitespace",
		Description:   "AttributeValue in subject RelativeDistinguishedName sequence should not have trailing whitespace",
		Providence:    "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SubjectDNTrailingSpace{}})
}
