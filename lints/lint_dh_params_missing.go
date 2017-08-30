// lint_dh_params_missing.go

package lints

import (
	"crypto/dsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dsaParamsMissing struct {
	// Internal data here
}

func (l *dsaParamsMissing) Initialize() error {
	return nil
}

func (l *dsaParamsMissing) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaParamsMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	params := c.PublicKey.(*dsa.PublicKey).Parameters
	if params.P.BitLen() == 0 || params.Q.BitLen() == 0 || params.G.BitLen() == 0 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_params_missing",
		Description:   "DSA: Certificates MUST include all domain parameters",
		Source:        "BRs: 6.1.6",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &dsaParamsMissing{},
	})
}
