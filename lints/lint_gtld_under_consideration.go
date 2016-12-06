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

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
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
		return ResultStruct{Result: Warn, Details: string("CommonName " + c.Subject.CommonName + " contains invalid gtld")}, nil
	}
	for _, dnsname := range c.DNSNames {
		if util.IsValidGTLD(dnsname) == 0 {
			return ResultStruct{Result: Warn, Details: string("DNSName " + dnsname + " contains invalid gtld")}, nil
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "gtld_under_consideration",
		Description:   "CAs SHOULD NOT issue Certificates containing a new gTLD under consideration by ICANN.",
		Providence:    "CAB: 4.2.2",
		EffectiveDate: util.CABV113Date,
		Test:          &gtldUnderConsideration{}})
}
