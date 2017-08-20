// lint_sub_ca_no_ip_name_contstraints.go
/******************************************************************************************************************************
If the Subordinate CA Certificate is not allowed to issue certificates with an iPAddress, then the Subordinate
CA Certificate MUST specify the entire IPv4 and IPv6 address ranges in excludedSubtrees. The Subordinate
CA Certificate MUST include within excludedSubtrees an iPAddress GeneralName of 8 zero octets (covering
the IPv4 address range of 0.0.0.0/0). The Subordinate CA Certificate MUST also include within
excludedSubtrees an iPAddress GeneralName of 32 zero octets (covering the IPv6 address range of ::0/0).
Otherwise, the Subordinate CA Certificate MUST include at least one iPAddress in permittedSubtrees.
******************************************************************************************************************************/

package lints

import (
	"bytes"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"net"
)

type subCaBadIPConstraint struct {
	// Internal data here
}

func (l *subCaBadIPConstraint) Initialize() error {
	return nil
}

func (l *subCaBadIPConstraint) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.NameConstOID)
}

func (l *subCaBadIPConstraint) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if len(c.PermittedIPAddresses) == 0 {
		v4 := false
		v6 := false
		for _, ips := range c.ExcludedIPAddresses {
			if bytes.Equal(ips.Data.IP, net.IP{0, 0, 0, 0}) && (bytes.Equal(ips.Data.Mask, net.IPMask{0, 0, 0, 0})) {
				v4 = true
				continue
			}
			if bytes.Equal(ips.Data.IP, net.IPv6zero) && bytes.Equal(ips.Data.Mask, net.IPMask{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) {
				v6 = true
				continue
			}
		}
		if v4 && v6 {
			return ResultStruct{Result: Pass}, nil
		} else {
			return ResultStruct{Result: Error}, nil
		}
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_no_ip_name_constraints",
		Description:   "Subordinate CA certs must include in the name contraints extension either permitted IP ranges or prohibit all IP addresses",
		Provenance:    "CAB: 7.1.5",
		EffectiveDate: util.CABV116Date,
		Test:          &subCaBadIPConstraint{},
	})
}
