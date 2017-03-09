/* zlint.go
 * Used to check parsed info from certificate for compliance
 */

package zlint

import (
	"encoding/base64"
	"errors"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lints"
)

//Calls all other checks on parsed certs
func ParsedTestHandler(cert *x509.Certificate) (map[string]lints.FinalResult, error) {
	if cert == nil {
		return nil, errors.New("zlint: nil pointer passed in, no data returned")
	}
	//run all tests
	var out map[string]lints.FinalResult = make(map[string]lints.FinalResult)
	for _, l := range lints.Lints {
		result, err := l.ExecuteTest(cert)
		finalResult := lints.FinalResult{}
		finalResult.Result = result.Result
		if err != nil {
			return out, err
		}
		out[l.Name] = finalResult
	}

	return out, nil
}

//expects Base64 encoded ASN.1 DER string, wrapper for ParsedTestHandler
func Lint64(certIn string) (map[string]lints.FinalResult, error) {

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
