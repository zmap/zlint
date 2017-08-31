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

const Version = 3

// ZLint contains the output of running all lints against a single certificate.
type ZLint struct {
	ZLintVersion    int64                         `json:"version"`
	Timestamp       int64                         `json:"timestamp"`
	ZLint           map[string]lints.ResultStruct `json:"lints"`
	NoticesPresent  bool                          `json:"notices_present"`
	WarningsPresent bool                          `json:"warnings_present"`
	ErrorsPresent   bool                          `json:"errors_present"`
	FatalsPresent   bool                          `json:"fatals_present"`
}

func (z *ZLint) execute(cert *x509.Certificate) {
	z.ZLint = make(map[string]lints.ResultStruct, len(lints.Lints))
	for name, l := range lints.Lints {
		res := l.Execute(cert)
		z.ZLint[name] = res
		z.updateErrorStatePresent(&res)
	}
}

func (z *ZLint) updateErrorStatePresent(result *lints.ResultStruct) {
	switch result.Result {
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
func LintCertificate(c *x509.Certificate) *ZLint {
	// Instead of panicing on nil certificate, just returns nil and let the client
	// panic when accessing ZLint, if they're into panicing.
	if c == nil {
		return nil
	}

	// Run all tests
	res := new(ZLint)
	res.execute(c)
	res.ZLintVersion = Version
	res.Timestamp = time.Now().Unix()
	return res
}
