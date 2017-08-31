package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type caIsCA struct {
	// Internal data here
}

type basicConstraints struct {
	IsCA       bool `asn1:"optional"`
	MaxPathLen int  `asn1:"optional,default:-1"`
}

func (l *caIsCA) Initialize() error {
	return nil
}

func (l *caIsCA) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsExtInCert(c, util.KeyUsageOID) && c.KeyUsage&x509.KeyUsageCertSign != 0 && util.IsExtInCert(c, util.BasicConstOID)
}

func (l *caIsCA) RunTest(c *x509.Certificate) (ResultStruct, error) {
	e := util.GetExtFromCert(c, util.BasicConstOID)
	var constraints basicConstraints
	_, err := asn1.Unmarshal(e.Value, &constraints)
	if err != nil {
		return ResultStruct{Result: Fatal}, nil
	}
	if constraints.IsCA == true {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_is_ca",
		Description:   "Root and Sub CA Certificate: The CA field MUST be set to true.",
		Source:        "BRs: 7.1.2.1, BRs: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &caIsCA{},
	})
}
