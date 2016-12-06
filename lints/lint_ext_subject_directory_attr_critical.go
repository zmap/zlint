// lint_ext_subject_directory_attr_critical.go
/************************************************
RFC 5280: 4.2.1.8
The subject directory attributes extension is used to convey
   identification attributes (e.g., nationality) of the subject.  The
   extension is defined as a sequence of one or more attributes.
   Conforming CAs MUST mark this extension as non-critical.
************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type subDirAttrCrit struct {
	// Internal data here
}

func (l *subDirAttrCrit) Initialize() error {
	return nil
}

func (l *subDirAttrCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectDirAttrOID)
}

func (l *subDirAttrCrit) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if e := util.GetExtFromCert(c, util.SubjectDirAttrOID); e.Critical {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "ext_subject_directory_attr_critical",
		Description:   "Conforming CAs must mark the Subject Directory Attributes extension as not critical",
		Providence:    "RFC 5280: 4.2.1.8",
		EffectiveDate: util.RFC2459Date,
		Test:          &subDirAttrCrit{}})
}
