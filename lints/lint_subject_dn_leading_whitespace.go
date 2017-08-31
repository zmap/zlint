// lint_subject_dn_leading_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubjectDNLeadingSpace struct{}

func (l *SubjectDNLeadingSpace) Initialize() error {
	return nil
}

func (l *SubjectDNLeadingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectDNLeadingSpace) Execute(c *x509.Certificate) ResultStruct {
	leading, _, err := util.CheckRDNSequenceWhiteSpace(c.RawSubject)
	if err != nil {
		return ResultStruct{Result: Fatal}
	}
	if leading {
		return ResultStruct{Result: Warn}
	}
	return ResultStruct{Result: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_subject_dn_leading_whitespace",
		Description:   "AttributeValue in subject RelativeDistinguishedName sequence SHOULD NOT have leading whitespace",
		Source:        "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &SubjectDNLeadingSpace{},
	})
}
