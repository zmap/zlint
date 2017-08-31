// lint_subject_dn_trailing_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubjectDNTrailingSpace struct{}

func (l *SubjectDNTrailingSpace) Initialize() error {
	return nil
}

func (l *SubjectDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectDNTrailingSpace) Execute(c *x509.Certificate) ResultStruct {
	_, trailing, err := util.CheckRDNSequenceWhiteSpace(c.RawSubject)
	if err != nil {
		return ResultStruct{Result: Fatal}
	}
	if trailing {
		return ResultStruct{Result: Warn}
	}
	return ResultStruct{Result: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_subject_dn_trailing_whitespace",
		Description:   "AttributeValue in subject RelativeDistinguishedName sequence SHOULD NOT have trailing whitespace",
		Source:        "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &SubjectDNTrailingSpace{},
	})
}
