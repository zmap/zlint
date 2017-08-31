// lint_sub_ca_eku_critical.go
/************************************************
BRs: 7.1.2.2g extkeyUsage (optional)
For Subordinate CA Certificates to be Technically constrained in line with section 7.1.5, then either the value
id‐kp‐serverAuth [RFC5280] or id‐kp‐clientAuth [RFC5280] or both values MUST be present**.
Other values MAY be present.
If present, this extension SHOULD be marked non‐critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCAEKUCrit struct {
	// Internal data here
}

func (l *subCAEKUCrit) Initialize() error {
	return nil
}

func (l *subCAEKUCrit) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsSubCA(c) && util.IsExtInCert(c, util.EkuSynOid)
}

func (l *subCAEKUCrit) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if e := util.GetExtFromCert(c, util.EkuSynOid); e.Critical {
		return ResultStruct{Result: Warn}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_ca_eku_critical",
		Description:   "Subordinate CA certificate extkeyUsage extension should be marked non-critical if present",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.CABV116Date,
		Test:          &subCAEKUCrit{},
	})
}
