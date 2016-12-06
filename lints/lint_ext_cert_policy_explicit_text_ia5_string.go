// lint_ext_cert_policy_explicit_text_ia5_string.go
/********************************************************************

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
********************************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type explicitTextIa5 struct {
	// Internal data here
}

func (l *explicitTextIa5) Initialize() error {
	return nil
}

func (l *explicitTextIa5) CheckApplies(c *x509.Certificate) bool {
	for _, text := range c.ExplicitTexts {
		if text != nil {
			return true
		}
	}
	return false
}

func (l *explicitTextIa5) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, firstLvl := range c.ExplicitTexts {
		for _, text := range firstLvl {
			if text.Tag == 22 {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "ext_cert_policy_explicit_text_ia5_string",
		Description:   "Compliant certificates must not encode explicitTest as IA5String",
		Providence:    "RFC 6818: 3",
		EffectiveDate: util.RFC6818Date,
		Test:          &explicitTextIa5{}})
}
