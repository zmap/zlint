// lint_ext_ian_rfc822_format_invalid.go
/************************************************************************
RFC 5280: 4.2.1.6
 When the issuerAltName extension contains an Internet mail address,
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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type IANEmail struct {
	// Internal data here
}

func (l *IANEmail) Initialize() error {
	return nil
}

func (l *IANEmail) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANEmail) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, str := range c.IANEmailAddresses {
		if str == "" {
			continue
		}
		if strings.Contains(str, " ") {
			return ResultStruct{Result: Error}, nil
		} else if str[0] == '<' || str[len(str)-1] == ')' {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_ian_rfc822_format_invalid",
		Description:   "Email must not be surrounded with `<>`, and there MUST NOT be trailing comments in `()`",
		Source:        "RFC 5280: 4.2.1.7",
		EffectiveDate: util.RFC2459Date,
		Test:          &IANEmail{},
	})
}
