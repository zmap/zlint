// lint_ext_cert_policy_explicit_text_not_nfc.go
/************************************************
  When the UTF8String encoding is used, all character sequences SHOULD be
  normalized according to Unicode normalization form C (NFC) [NFC].
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/text/unicode/norm"
)

type ExtCertPolicyExplicitTextNotNFC struct{}

func (l *ExtCertPolicyExplicitTextNotNFC) Initialize() error {
	return nil
}

func (l *ExtCertPolicyExplicitTextNotNFC) CheckApplies(c *x509.Certificate) bool {
	for _, text := range c.ExplicitTexts {
		if text != nil {
			return true
		}
	}
	return false
}

func (l *ExtCertPolicyExplicitTextNotNFC) Execute(c *x509.Certificate) ResultStruct {
	for _, firstLvl := range c.ExplicitTexts {
		for _, text := range firstLvl {
			if text.Tag == 12 || text.Tag == 30 {
				if !norm.NFC.IsNormal(text.Bytes) {
					return ResultStruct{Result: Warn}
				}
			}
		}
	}
	return ResultStruct{Result: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_cert_policy_explicit_text_not_nfc",
		Description:   "When utf8string or bmpstring encoding is used for explicitText field in certificate policy, it SHOULD be normalized by NFC format",
		Source:        "Fill this in...",
		EffectiveDate: util.RFC6818Date,
		Lint:          &ExtCertPolicyExplicitTextNotNFC{},
	})
}
