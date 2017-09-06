// lint_wrong_time_format_pre2050.go
/*********************************************************************
CAs conforming to this profile MUST always encode certificate
validity dates through the year 2049 as UTCTime; certificate validity
dates in 2050 or later MUST be encoded as GeneralizedTime.
Conforming applications MUST be able to process validity dates that
are encoded in either UTCTime or GeneralizedTime.
*********************************************************************/

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"time"
)

type generalizedPre2050 struct{}

func (l *generalizedPre2050) Initialize() error {
	return nil
}

func (l *generalizedPre2050) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *generalizedPre2050) Execute(c *x509.Certificate) *LintResult {
	date1, date2 := util.GetTimes(c)
	var t time.Time
	type1, type2 := util.FindTimeType(date1, date2)
	if type1 == 24 {
		temp, err := asn1.Marshal(date1)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		_, err = asn1.Unmarshal(temp, &t)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		if t.Before(util.GeneralizedDate) {
			return &LintResult{Status: Error}
		}
	}
	if type2 == 24 {
		temp, err := asn1.Marshal(date2)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		_, err = asn1.Unmarshal(temp, &t)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		if t.Before(util.GeneralizedDate) {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_wrong_time_format_pre2050",
		Description:    "Certificates valid through the year 2049 MUST be encoded in UTC time",
		ReadableSource: "RFC 5280: 4.1.2.5",
		Source:         RFC5280,
		EffectiveDate:  util.RFC2459Date,
		Lint:           &generalizedPre2050{},
	})
}
