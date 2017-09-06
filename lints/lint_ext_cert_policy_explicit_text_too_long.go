// lint_ext_cert_policy_explicit_text_too_long.go
/*******************************************************************
An explicitText field includes the textual statement directly in
the certificate.  The explicitText field is a string with a
maximum size of 200 characters.  Conforming CAs SHOULD use the
UTF8String encoding for explicitText.  VisibleString or BMPString
are acceptable but less preferred alternatives.  Conforming CAs
MUST NOT encode explicitText as IA5String.  The explicitText string
SHOULD NOT include any control characters (e.g., U+0000 to U+001F
and U+007F to U+009F).  When the UTF8String or BMPString encoding
is used, all character sequences SHOULD be normalized according
to Unicode normalization form C (NFC) [NFC].
*******************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type explicitTextTooLong struct{}

func (l *explicitTextTooLong) Initialize() error {
	return nil
}

func (l *explicitTextTooLong) CheckApplies(c *x509.Certificate) bool {
	for _, text := range c.ExplicitTexts {
		if text != nil {
			return true
		}
	}
	return false
}

func (l *explicitTextTooLong) Execute(c *x509.Certificate) *LintResult {
	for _, firstLvl := range c.ExplicitTexts {
		for _, text := range firstLvl {
			if len(text.Bytes) > 200 {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_ext_cert_policy_explicit_text_too_long",
		Description:    "Explicit text has a maximum size of 200 characters",
		ReadableSource: "RFC 6818: 3",
		Source:         RFC5280,
		EffectiveDate:  util.RFC6818Date,
		Lint:           &explicitTextTooLong{},
	})
}
