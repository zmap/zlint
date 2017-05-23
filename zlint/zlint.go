/* zlint.go
 * Used to check parsed info from certificate for compliance
 */

package zlint

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lints"
	"time"
	"encoding/json"
	"bytes"
)

type PrettyOutput struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Providence string `json:"providence"`
}

//Pretty Print lint outputs
func PrettyPrintZLint() {
	for _, l := range lints.Lints {
		p := PrettyOutput{}
		p.Name = l.Name
		p.Description = l.Description
		p.Providence = l.Providence
		out, _ := json.Marshal(p)
		out = bytes.Replace(out, []byte("\\u003c"), []byte("<"), -1)
		out = bytes.Replace(out, []byte("\\u003e"), []byte(">"), -1)
		out = bytes.Replace(out, []byte("\\u0026"), []byte("&"), -1)
		fmt.Println(string(out))
	}
}

//Calls all other checks on parsed certs producing version
func ZLintResultTestHandler(cert *x509.Certificate) *lints.ZLintResult {
	if cert == nil {
		return nil
	}
	//run all tests
	ZLintResult := lints.ZLintResult{}
	ZLintReport := lints.LintReport{}
	ZLintResult.ZLint = &ZLintReport
	ZLintResult.Execute(cert)
	ZLintResult.ZLintVersion = lints.ZLintVersion
	ZLintResult.Timestamp = time.Now().Unix()
	return &ZLintResult
}

//Calls all other checks on parsed certs
func ParsedTestHandler(cert *x509.Certificate) (map[string]string, error) {
	if cert == nil {
		return nil, errors.New("zlint: nil pointer passed in, no data returned")
	}
	//run all tests
	var out map[string]string = make(map[string]string)
	for _, l := range lints.Lints {
		result, err := l.ExecuteTest(cert)
		if err != nil {
			return out, err
		}
		out[l.Name] = lints.EnumToString(result.Result)
	}

	return out, nil
}

//expects Base64 encoded ASN.1 DER string, wrapper for ParsedTestHandler
func Lint64(certIn string) (map[string]string, error) {

	der, err := base64.StdEncoding.DecodeString(certIn) //decode string
	if err != nil {
		return nil, err //parsing error, no tests performed
	}

	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err //parsing error, no tests performed
	}

	return ParsedTestHandler(cert) //return available reports & error from main testing function
}
