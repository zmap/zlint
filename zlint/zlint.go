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
func ParsedTestHandler(cert *x509.Certificate, m map[string]int) (map[string]lints.FinalResult, error) {
	if cert == nil {
		return nil, errors.New("zlint: nil pointer passed in, no data returned")
	}
	//run all tests
	var out map[string]lints.FinalResult = make(map[string]lints.FinalResult)
	for _, l := range lints.Lints {
		result, err := l.ExecuteTest(cert)
		if err != nil {
			return out, err
		}
		out[l.Name] = lints.FinalResult{Result: enumToString(result.Result), Details: result.Details}
		if result.Result == lints.Warn || result.Result == lints.Error {
			m[l.Name]++
		}
	}

	return out, nil
}

//expects Base64 encoded ASN.1 DER string, wrapper for ParsedTestHandler
func Lint64(certIn string, m map[string]int) (map[string]lints.FinalResult, error) {

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

func enumToString(e lints.ResultEnum) string {
	switch e {
	case lints.NA:
		return "NA"
	case lints.NE:
		return "NE"
	case lints.Pass:
		return "pass"
	case lints.Info:
		return "info"
	case lints.Warn:
		return "warn"
	case lints.Error:
		return "error"
	case lints.Fatal:
		return "fatal"
	default:
		return ""
	}
}
