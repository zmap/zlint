// lint_ext_cert_policy_explicit_text_includes_control.go
/*********************************************************************
An explicitText field includes the textual statement directly in
the certificate.  The explicitText field is a string with a
maximum size of 200 characters.  Conforming CAs SHOULD use the
UTF8String encoding for explicitText, but MAY use IA5String.
Conforming CAs MUST NOT encode explicitText as VisibleString or
BMPString.  The explicitText string SHOULD NOT include any control
characters (e.g., U+0000 to U+001F and U+007F to U+009F).  When
the UTF8String encoding is used, all character sequences SHOULD be
normalized according to Unicode normalization form C (NFC) [NFC].
*********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type controlChar struct {
	// Internal data here
}

func (l *controlChar) Initialize() error {
	return nil
}

func (l *controlChar) CheckApplies(c *x509.Certificate) bool {
	for _, text := range c.ExplicitTexts {
		if text != nil {
			return true
		}
	}
	return false
}

func (l *controlChar) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, firstLvl := range c.ExplicitTexts {
		for _, text := range firstLvl {
			if text.Tag == 12 {
				for i := 0; i < len(text.Bytes); i++ {
					if text.Bytes[i]&0x80 == 0 {
						if text.Bytes[i] < 0x20 || text.Bytes[i] == 0x7f {
							return ResultStruct{Result: Warn}, nil
						}
					} else if text.Bytes[i]&0x20 == 0 {
						if text.Bytes[i] == 0xc2 && text.Bytes[i+1] >= 0x80 && text.Bytes[i+1] <= 0x9f {
							return ResultStruct{Result: Warn}, nil
						}
						i += 1
					} else if text.Bytes[i]&0x10 == 0 {
						i += 2
					} else if text.Bytes[i]&0x08 == 0 {
						i += 3
					} else if text.Bytes[i]&0x04 == 0 {
						i += 4
					} else if text.Bytes[i]&0x02 == 0 {
						i += 5
					}
				}
			}
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_cert_policy_explicit_text_includes_control",
		Description:   "Explicit text should not include any control charaters",
		Source:        "RFC 6818: 3",
		EffectiveDate: util.RFC6818Date,
		Test:          &controlChar{},
	})
}
