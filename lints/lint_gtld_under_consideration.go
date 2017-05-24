// lint_gtld_under_consideration.go
/************************************************
CAs SHOULD NOT issue Certificates containing a new
gTLD under consideration by ICANN. Prior to issuing a
Certificate containing an Internal Name with a gTLD
that ICANN has announced as under consideration to
make operational, the CA MUST provide a warning to
the applicant that the gTLD may soon become
resolvable and that, at that time, the CA will revoke
the Certificate unless the applicant promptly registers the
domain name.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type gtldUnderConsideration struct {
	// Internal data here
}

func (l *gtldUnderConsideration) Initialize() error {
	return nil
}

func (l *gtldUnderConsideration) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *gtldUnderConsideration) RunTest(c *x509.Certificate) (ResultStruct, error) {
	//Need to run the GTLD test for both Subject.Common_Name and DNSNames(from SAN)
	if util.IsValidGTLD(c.Subject.CommonName) == 0 {
		return ResultStruct{Result: Warn}, nil
	}
	for _, dnsname := range c.DNSNames {
		if util.IsValidGTLD(dnsname) == 0 {
			return ResultStruct{Result: Warn}, nil
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_gtld_under_consideration",
		Description:   "CAs SHOULD NOT issue certificates containing a new gTLD under consideration by ICANN",
		Providence:    "CAB: 4.2.2",
		EffectiveDate: util.CABV113Date,
		Test:          &gtldUnderConsideration{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.WGtldUnderConsideration = result },
	})
}
