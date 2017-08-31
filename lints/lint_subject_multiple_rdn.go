// lint_subject_multiple_attr_in_rdn.go

package lints

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
)

type SubjectRDNHasMultipleAttribute struct{}

func (l *SubjectRDNHasMultipleAttribute) Initialize() error {
	return nil
}

func (l *SubjectRDNHasMultipleAttribute) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectRDNHasMultipleAttribute) Execute(c *x509.Certificate) ResultStruct {
	var subject pkix.RDNSequence
	if _, err := asn1.Unmarshal(c.RawSubject, &subject); err != nil {
		return ResultStruct{Result: Fatal}
	}
	for _, rdn := range subject {
		if len(rdn) > 1 {
			return ResultStruct{Result: Warn}
		}
	}
	return ResultStruct{Result: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_multiple_subject_rdn",
		Description:   "Certificates should not have multiple attributes in a single RDN (subject)",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &SubjectRDNHasMultipleAttribute{},
	})
}
