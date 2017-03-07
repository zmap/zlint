// testingUtil.go
// Contains resources necessary to the Unit Test Cases

package lints

import (
	"encoding/pem"
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"io/ioutil"
	"strings"
)

func ReadCertificate(inPath string) *x509.Certificate {
	// All of this can be encapsulated in a function
	data, err := ioutil.ReadFile(inPath)
	if err != nil {
		//read failure, die horribly here
		fmt.Println(err)
		panic("File read failed!")
	}
	var textData string = string(data)
	if strings.Contains(textData, "-BEGIN CERTIFICATE-") {
		block, _ := pem.Decode(data)
		if block == nil {
			panic("PEM decode failed!")
		}
		data = block.Bytes
	}
	theCert, err := x509.ParseCertificate(data)
	if err != nil {
		//die horribly here
		fmt.Println(err)
		return nil
	}
	return theCert
}
