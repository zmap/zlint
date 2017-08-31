/* zlint.go
 * Used to check parsed info from certificate for compliance
 */

package zlint

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lints"
)

type LintDescription struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Source      string `json:"source"`
}

// EncodeLintDescriptionsToJSON outputs a description of the lint as JSON.
func EncodeLintDescriptionsToJSON(w io.Writer) {
	for _, l := range lints.Lints {
		p := LintDescription{}
		p.Name = l.Name
		p.Description = l.Description
		p.Source = l.Source

		buffer := new(bytes.Buffer)
		enc := json.NewEncoder(buffer)
		enc.SetEscapeHTML(false)
		enc.Encode(p)
		w.Write(buffer.Bytes())
	}
}

// LintCertificate runs all registered lints on c, producing a ZLintResult.
func LintCertificate(c *x509.Certificate) *lints.ZLintResult {
	if c == nil {
		return nil
	}
	//run all tests
	ZLintResult := lints.ZLintResult{}
	ZLintResult.Execute(c)
	ZLintResult.ZLintVersion = lints.ZLintVersion
	ZLintResult.Timestamp = time.Now().Unix()
	return &ZLintResult
}
