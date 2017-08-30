// lint_generalized_time_not_in_zulu.go
/********************************************************************
4.1.2.5.2.  GeneralizedTime
The generalized time type, GeneralizedTime, is a standard ASN.1 type
for variable precision representation of time.  Optionally, the
GeneralizedTime field can include a representation of the time
differential between local and Greenwich Mean Time.

For the purposes of this profile, GeneralizedTime values MUST be
expressed in Greenwich Mean Time (Zulu) and MUST include seconds
(i.e., times are YYYYMMDDHHMMSSZ), even where the number of seconds
is zero.  GeneralizedTime values MUST NOT include fractional seconds.
********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type generalizedNotZulu struct {
	date1Gen bool
	date2Gen bool
}

func (l *generalizedNotZulu) Initialize() error {
	return nil
}

func (l *generalizedNotZulu) CheckApplies(c *x509.Certificate) bool {
	firstDate, secondDate := util.GetTimes(c)
	beforeTag, afterTag := util.FindTimeType(firstDate, secondDate)
	l.date1Gen = beforeTag == 24
	l.date2Gen = afterTag == 24
	return l.date1Gen || l.date2Gen
}

func (l *generalizedNotZulu) RunTest(c *x509.Certificate) (ResultStruct, error) {
	date1, date2 := util.GetTimes(c)
	if l.date1Gen {
		if date1.Bytes[len(date1.Bytes)-1] != 'Z' {
			return ResultStruct{Result: Error}, nil
		}
	}
	if l.date2Gen {
		if date2.Bytes[len(date2.Bytes)-1] != 'Z' {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_generalized_time_not_in_zulu",
		Description:   "Generalized time values MUST be expressed in Greenwich Mean Time (Zulu)",
		Source:        "RFC 5280: 4.1.2.5.2",
		EffectiveDate: util.RFC2459Date,
		Test:          &generalizedNotZulu{},
	})
}
