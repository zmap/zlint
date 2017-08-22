package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCAEKUValidFields struct {
	// Internal data here
}

func (l *subCAEKUValidFields) Initialize() error {
	return nil
}

func (l *subCAEKUValidFields) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.EkuSynOid)
}

func (l *subCAEKUValidFields) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, ekuValue := range c.ExtKeyUsage {
		if ekuValue == x509.ExtKeyUsageServerAuth ||
			ekuValue == x509.ExtKeyUsageClientAuth {
			continue
		} else {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_eku_valid_fields",
		Description:   "Subordinate CA extkeyUsage, either id-kp-serverAuth or id-kp-clientAuth or both values MUST be present.",
		Provenance:    "BRs: 7.1.2.2",
		EffectiveDate: util.CABV116Date,
		Test:          &subCAEKUValidFields{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ESubCaEkuValidFields = result },
	})
}
