// lint_ext_cert_policy_contains_noticeref.go
/********************************************************************
The user notice has two optional fields: the noticeRef field and the
explicitText field. Conforming CAs SHOULD NOT use the noticeRef
option.
********************************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type noticeRefPres struct {
	// Internal data here
}

func (l *noticeRefPres) Initialize() error {
	return nil
}

func (l *noticeRefPres) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CertPolicyOID)
}

func (l *noticeRefPres) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, firstLvl := range c.NoticeRefNumbers {
		for _, number := range firstLvl {
			if number != nil {
				return ResultStruct{Result: Warn}, nil
			}
		}
	}
	for _, firstLvl := range c.NoticeRefOrgnization {
		for _, org := range firstLvl {
			if len(org.Bytes) != 0 {
				return ResultStruct{Result: Warn}, nil
			}
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "ext_cert_policy_contains_noticeref",
		Description:   "Compliant certificates should not use the noticeRef option",
		Providence:    "RFC 5280: 4.2.1.4",
		EffectiveDate: util.RFC5280Date,
		Test:          &noticeRefPres{}})
}
