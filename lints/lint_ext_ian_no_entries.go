// lint_ext_ian_no_entries.go
/**********************************************************************
RFC 5280: 4.2.1.7
If the issuerAltName extension is present, the sequence MUST contain
   at least one entry.  Unlike the subject field, conforming CAs MUST
   NOT issue certificates with subjectAltNames containing empty
   GeneralName fields.  For example, an rfc822Name is represented as an
   IA5String.  While an empty string is a valid IA5String, such an
   rfc822Name is not permitted by this profile.  The behavior of clients
   that encounter such a certificate when processing a certification
   path is not defined by this profile.
***********************************************************************/

package lints

import (
	"github.com/zmap/zgrab/ztools/x509"
	"github.com/zmap/zlint/util"
)

type ianNoEntry struct {
	// Internal data here
}

func (l *ianNoEntry) Initialize() error {
	return nil
}

func (l *ianNoEntry) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerANOID)
}

func (l *ianNoEntry) RunTest(c *x509.Certificate) (ResultStruct, error) {
	ian := util.GetExtFromCert(c, util.IssuerANOID)
	if (ian.Value)[1] == 0 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "ext_ian_no_entries",
		Description:   "if present, the ian extension must contain at least one entry",
		Providence:    "RFC 5280: 4.2.1.7",
		EffectiveDate: util.RFC2459Date,
		Test:          &ianNoEntry{}})
}
