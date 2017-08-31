/* z.go
 * Used to check parsed info from certificate for compliance
 */

package zlint

import (
	"encoding/json"
	"io"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lints"
)

const Version int64 = 3

// ResultSet contains the output of running all lints against a single certificate.
type ResultSet struct {
	Version         int64                        `json:"version"`
	Timestamp       int64                        `json:"timestamp"`
	Results         map[string]*lints.LintResult `json:"lints"`
	NoticesPresent  bool                         `json:"notices_present"`
	WarningsPresent bool                         `json:"warnings_present"`
	ErrorsPresent   bool                         `json:"errors_present"`
	FatalsPresent   bool                         `json:"fatals_present"`
}

func (z *ResultSet) execute(cert *x509.Certificate) {
	z.Results = make(map[string]*lints.LintResult, len(lints.Lints))
	for name, l := range lints.Lints {
		res := l.Execute(cert)
		z.Results[name] = res
		z.updateErrorStatePresent(res)
	}
}

func (z *ResultSet) updateErrorStatePresent(result *lints.LintResult) {
	switch result.Status {
	case lints.Notice:
		z.NoticesPresent = true
	case lints.Warn:
		z.WarningsPresent = true
	case lints.Error:
		z.ErrorsPresent = true
	case lints.Fatal:
		z.FatalsPresent = true
	}
}

// EncodeLintDescriptionsToJSON outputs a description of each lint as JSON
// object, one object per line.
func EncodeLintDescriptionsToJSON(w io.Writer) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	for _, lint := range lints.Lints {
		enc.Encode(lint)
	}
}

// LintCertificate runs all registered lints on c, producing a ZLint.
func LintCertificate(c *x509.Certificate) *ResultSet {
	// Instead of panicing on nil certificate, just returns nil and let the client
	// panic when accessing ZLint, if they're into panicing.
	if c == nil {
		return nil
	}

	// Run all tests
	res := new(ResultSet)
	res.execute(c)
	res.Version = Version
	res.Timestamp = time.Now().Unix()
	return res
}
