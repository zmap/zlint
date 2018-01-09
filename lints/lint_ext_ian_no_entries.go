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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type IANNoEntry struct{}

func (l *IANNoEntry) Initialize() error {
	return nil
}

func (l *IANNoEntry) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANNoEntry) Execute(c *x509.Certificate) *LintResult {
	ian := util.GetExtFromCert(c, util.IssuerAlternateNameOID)
	if util.IsEmptyASN1Sequence(ian.Value) {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_ian_no_entries",
		Description:   "If present, the IAN extension must contain at least one entry",
		Citation:      "RFC 5280: 4.2.1.7",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &IANNoEntry{},
	})
}
