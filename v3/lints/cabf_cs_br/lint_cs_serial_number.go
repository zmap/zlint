package cabf_cs_br

import (
	"github.com/zmap/zcrypto/x509"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

/*7.1.1 Serial Number
CAs SHALL generate non‑sequential Certificate serial numbers greater than
zero (0) containing at least 64 bits of output from a CSPRNG.
*/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_cs_serial_number",
			Description:   "CAs SHALL generate non-sequential Certificate serial numbers greater than zero (0) containing at least 64 bits of output from a CSPRNG",
			Citation:      "CABF CS BRs 7.1.1",
			Source:        lint.CABFCSBaselineRequirements,
			EffectiveDate: util.CABF_CS_BRs_1_2_Date,
		},
		Lint: NewCsSerialNumber,
	})
}

type csSerialNumber struct{}

func NewCsSerialNumber() lint.CertificateLintInterface {
	return &csSerialNumber{}
}

func (l *csSerialNumber) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *csSerialNumber) Execute(c *x509.Certificate) *lint.LintResult {
	if c.SerialNumber == nil || c.SerialNumber.Sign() <= 0 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "certificate serial number must be greater than zero",
		}
	}

	if c.SerialNumber.BitLen() < 64 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "certificate serial number must contain at least 64 bits of CSPRNG output",
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
