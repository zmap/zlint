package cabf_cs_br

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

/*7.1.2.3 a. authorityInformationAccess
It MUST contain the HTTP URL of the Issuing CA's certificate
(accessMethod = 1.3.6.1.5.5.7.48.2).*/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_cs_aia_missing_ca_issuers_http_url",
			Description:   "The authorityInformationAccess extension MUST contain the HTTP URL of the Issuing CA's certificate (id-ad-caIssuers).",
			Citation:      "CABF CS BRs 7.1.2.3.a",
			Source:        lint.CABFCSBaselineRequirements,
			EffectiveDate: util.CABF_CS_BRs_1_2_Date,
		},
		Lint: NewCsAiaMissingCaIssuersHttpUrl,
	})
}

type csAiaMissingCaIssuersHttpUrl struct{}

func NewCsAiaMissingCaIssuersHttpUrl() lint.LintInterface {
	return &csAiaMissingCaIssuersHttpUrl{}
}

func (l *csAiaMissingCaIssuersHttpUrl) CheckApplies(c *x509.Certificate) bool {
	return (util.IsSubscriberCert(c) || util.IsSubCA(c))
}

func (l *csAiaMissingCaIssuersHttpUrl) Execute(c *x509.Certificate) *lint.LintResult {
	if len(c.IssuingCertificateURL) == 0 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "authorityInformationAccess MUST contain the HTTP URL of the Issuing CA's certificate (id-ad-caIssuers)."}
	}

	for _, uri := range c.IssuingCertificateURL {
		if strings.HasPrefix(uri, "http://") {
			return &lint.LintResult{Status: lint.Pass}
		}
	}
	return &lint.LintResult{
		Status:  lint.Error,
		Details: "authorityInformationAccess MUST contain the HTTP URL of the Issuing CA's certificate (id-ad-caIssuers)."}
}
