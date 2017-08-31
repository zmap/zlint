// lint_utc_time_not_in_zulu.go
/***********************************************************************
4.1.2.5.1.  UTCTime
   The universal time type, UTCTime, is a standard ASN.1 type intended
   for representation of dates and time.  UTCTime specifies the year
   through the two low-order digits and time is specified to the
   precision of one minute or one second.  UTCTime includes either Z
   (for Zulu, or Greenwich Mean Time) or a time differential.

   For the purposes of this profile, UTCTime values MUST be expressed in
   Greenwich Mean Time (Zulu) and MUST include seconds (i.e., times are
   YYMMDDHHMMSSZ), even where the number of seconds is zero.  Conforming
   systems MUST interpret the year field (YY) as follows:

      Where YY is greater than or equal to 50, the year SHALL be
      interpreted as 19YY; and

      Where YY is less than 50, the year SHALL be interpreted as 20YY.
***********************************************************************/
package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"time"
)

type utcTimeGMT struct {
	date1Utc bool
	date2Utc bool
}

func (l *utcTimeGMT) Initialize() error {
	return nil
}

func (l *utcTimeGMT) CheckApplies(c *x509.Certificate) bool {
	firstDate, secondDate := util.GetTimes(c)
	beforeTag, afterTag := util.FindTimeType(firstDate, secondDate)
	l.date1Utc = beforeTag == 23
	l.date2Utc = afterTag == 23
	return l.date1Utc || l.date2Utc
}

func (l *utcTimeGMT) RunTest(c *x509.Certificate) (ResultStruct, error) {
	var r ResultEnum
	if l.date1Utc {
		// UTC Tests on notBefore
		utcNotGmt(c.NotBefore, &r)
	}
	if l.date2Utc {
		// UTC Tests on NotAfter
		utcNotGmt(c.NotAfter, &r)
	}
	return ResultStruct{Result: r}, nil
}

func utcNotGmt(t time.Time, r *ResultEnum) {
	// If we already ran this test and it resulted in error, don't want to discard that
	// And now we use the afterBool to make sure we test the right time
	if *r == Error {
		return
	}
	if t.Location() != time.UTC {
		*r = Error
	} else {
		*r = Pass
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_utc_time_not_in_zulu",
		Description:   "UTCTime values MUST be expressed in Greenwich Mean Time (Zulu)",
		Source:        "RFC 5280: 4.1.2.5.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &utcTimeGMT{},
	})
}
