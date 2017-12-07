// lint_ext_san_rfc822_format_invalid.go
/************************************************************************
RFC 5280: 4.2.1.6
 When the subjectAltName extension contains an Internet mail address,
   the address MUST be stored in the rfc822Name.  The format of an
   rfc822Name is a "Mailbox" as defined in Section 4.1.2 of [RFC2821].
   A Mailbox has the form "Local-part@Domain".  Note that a Mailbox has
   no phrase (such as a common name) before it, has no comment (text
   surrounded in parentheses) after it, and is not surrounded by "<" and
   ">".  Rules for encoding Internet mail addresses that include
   internationalized domain names are specified in Section 7.5.
************************************************************************/

package lints

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type invalidEmail struct{}

func (l *invalidEmail) Initialize() error {
	return nil
}

func (l *invalidEmail) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *invalidEmail) Execute(c *x509.Certificate) *LintResult {
	for _, str := range c.EmailAddresses {
		if str == "" {
			continue
		}
		if strings.Contains(str, " ") {
			return &LintResult{Status: Error}
		} else if str[0] == '<' || str[len(str)-1] == ')' {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_rfc822_format_invalid",
		Description:   "Email MUST NOT be surrounded with `<>`, and there must be no trailing comments in `()`",
		Citation:      "RFC 5280: 4.2.1.6",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &invalidEmail{},
	})
}
