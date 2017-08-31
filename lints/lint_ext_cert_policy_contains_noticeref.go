// lint_ext_cert_policy_contains_noticeref.go
/********************************************************************
The user notice has two optional fields: the noticeRef field and the
explicitText field. Conforming CAs SHOULD NOT use the noticeRef
option.
********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type noticeRefPres struct{}

func (l *noticeRefPres) Initialize() error {
	return nil
}

func (l *noticeRefPres) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CertPolicyOID)
}

func (l *noticeRefPres) Execute(c *x509.Certificate) *LintResult {
	for _, firstLvl := range c.NoticeRefNumbers {
		for _, number := range firstLvl {
			if number != nil {
				return &LintResult{Status: Warn}
			}
		}
	}
	for _, firstLvl := range c.NoticeRefOrgnization {
		for _, org := range firstLvl {
			if len(org.Bytes) != 0 {
				return &LintResult{Status: Warn}
			}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_cert_policy_contains_noticeref",
		Description:   "Compliant certificates SHOULD NOT use the noticeRef option",
		Source:        "RFC 5280: 4.2.1.4",
		EffectiveDate: util.RFC5280Date,
		Lint:          &noticeRefPres{},
	})
}
