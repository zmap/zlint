/* zlint.go
 * Used to check parsed info from certificate for compliance
 */

package zlint

import (
	"encoding/base64"
	"errors"
	"github.com/zmap/zgrab/ztools/x509"
	"github.com/zmap/zlint/lints"
)

//Calls all other checks on parsed certs
func ParsedTestHandler(cert *x509.Certificate, m map[string]int) (map[string]lints.ResultStruct, error) {
	if cert == nil {
		return nil, errors.New("zlint: nil pointer passed in, no data returned")
	}
	//run all tests
	var out map[string]lints.ResultStruct = make(map[string]lints.ResultStruct)
	for _, l := range lints.Lints {
		result, err := l.ExecuteTest(cert)
		if err != nil {
			return out, err
		}
		out[l.Name] = result
		if result.Result == lints.Warn || result.Result == lints.Error {
			m[l.Name]++
		}
	}

	return out, nil
}

//expects Base64 encoded ASN.1 DER string, wrapper for ParsedTestHandler
func Lint64(certIn string, m map[string]int) (map[string]lints.ResultStruct, error) {

	der, err := base64.StdEncoding.DecodeString(certIn) //decode string
	if err != nil {
		return nil, err //parsing error, no tests performed
	}

	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err //parsing error, no tests performed
	}

	return ParsedTestHandler(cert, m) //return available reports & error from main testing function
}
