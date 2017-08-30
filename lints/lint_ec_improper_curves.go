// lint_ec_improper_curves.go
/************************************************
BRs: 6.1.5
Certificates MUST meet the following requirements for algorithm type and key size.
ECC Curve: NIST P-256, P-384, or P-521
************************************************/

package lints

import (
	"crypto/ecdsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ecImproperCurves struct {
	// Internal data here
}

func (l *ecImproperCurves) Initialize() error {
	return nil
}

func (l *ecImproperCurves) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.PublicKeyAlgorithm == x509.ECDSA
}

func (l *ecImproperCurves) RunTest(c *x509.Certificate) (ResultStruct, error) {
	/* Declare theKey to be a ECDSA Public Key */
	var theKey *ecdsa.PublicKey
	/* Need to do different things based on what c.PublicKey is */
	switch c.PublicKey.(type) {
	case *x509.AugmentedECDSA:
		temp := c.PublicKey.(*x509.AugmentedECDSA)
		theKey = temp.Pub
	case *ecdsa.PublicKey:
		theKey = c.PublicKey.(*ecdsa.PublicKey)
	}
	/* Now can actually check the params */
	theParams := theKey.Curve.Params()
	switch theParams.Name {
	case "P-256", "P-384", "P-521":
		return ResultStruct{Result: Pass}, nil
	default:
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:        "e_ec_improper_curves",
		Description: "Only one of NIST P‐256, P‐384, or P‐521 can be used",
		Source:      "BRs: 6.1.5",
		// Refer to BRs: 6.1.5, taking the statement "Before 31 Dec 2010" literally
		EffectiveDate: util.ZeroDate,
		Test:          &ecImproperCurves{},
	})
}
