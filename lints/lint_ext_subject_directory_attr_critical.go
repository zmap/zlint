/************************************************
RFC 5280: 4.2.1.8
The subject directory attributes extension is used to convey
   identification attributes (e.g., nationality) of the subject.  The
   extension is defined as a sequence of one or more attributes.
   Conforming CAs MUST mark this extension as non-critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subDirAttrCrit struct{}

func (l *subDirAttrCrit) Initialize() error {
	return nil
}

func (l *subDirAttrCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectDirAttrOID)
}

func (l *subDirAttrCrit) Execute(c *x509.Certificate) *LintResult {
	if e := util.GetExtFromCert(c, util.SubjectDirAttrOID); e.Critical {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_subject_directory_attr_critical",
		Description:   "Conforming CAs MUST mark the Subject Directory Attributes extension as not critical",
		Citation:      "RFC 5280: 4.2.1.8",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &subDirAttrCrit{},
	})
}
