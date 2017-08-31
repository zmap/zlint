package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCAEKUNameConstraints struct {
	// Internal data here
}

func (l *subCAEKUNameConstraints) Initialize() error {
	return nil
}

func (l *subCAEKUNameConstraints) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsSubCA(c) && util.IsExtInCert(c, util.EkuSynOid)
}

func includesNameConstraints(c *x509.Certificate) bool {
	if len(c.PermittedDNSNames) > 0 || len(c.ExcludedDNSNames) > 0 || len(c.PermittedIPAddresses) > 0 || len(c.ExcludedIPAddresses) > 0 || len(c.PermittedDirectoryNames) > 0 || len(c.ExcludedDirectoryNames) > 0 {
		return true
	} else {
		return false
	}
}

func (l *subCAEKUNameConstraints) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, eku := range c.ExtKeyUsage {
		if eku == x509.ExtKeyUsageServerAuth {
			if includesNameConstraints(c) {
				return ResultStruct{Result: Pass}, nil
			} else {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	return ResultStruct{Result: NA}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_eku_name_constraints",
		Description:   "Subordinate CA: If includes id-kp-serverAuth EKU, then it MUST include Name constraints w/ constraints on DNSName, IPAddress, and DirectoryName",
		Source:        "BRs: 7.1.5",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCAEKUNameConstraints{},
	})
}
