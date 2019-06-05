package util

import (
	"bytes"
	"encoding/asn1"
	"errors"

	"golang.org/x/crypto/cryptobyte"
	cryptobyte_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

// byte representation of pkix.AlgorithmIdentifier with OID rsaEncryption and Parameters asn1.NULL
var expectedAlgoIDBytes = []byte{0x30, 0x0d, 0x6, 0x9, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0xd, 0x1, 0x1, 0x1, 0x5, 0x0}

// CheckAlgorithmIDParamNotNULL parses an AlgorithmIdentifier with algorithm OID rsaEncryption to check the Param field is asn1.NULL
// Expects DER-encoded AlgorithmIdentifier including tag and length
func CheckAlgorithmIDParamNotNULL(algorithmIdentifier []byte) error {
	algorithmSequence := cryptobyte.String(algorithmIdentifier)

	// byte comparison of algorithm sequence and checking no trailing data is present
	var algorithmBytes []byte
	if algorithmSequence.ReadBytes(&algorithmBytes, len(expectedAlgoIDBytes)) {
		if bytes.Compare(algorithmBytes, expectedAlgoIDBytes) == 0 && algorithmSequence.Empty() {
			return nil
		}
	}

	// re-parse to get an error message detailing what did not match in the byte comparison
	algorithmSequence = cryptobyte.String(algorithmIdentifier)
	var algorithm cryptobyte.String
	if !algorithmSequence.ReadASN1(&algorithm, cryptobyte_asn1.SEQUENCE) {
		return errors.New("error reading algorithm")
	}

	encryptionOID := asn1.ObjectIdentifier{}
	if !algorithm.ReadASN1ObjectIdentifier(&encryptionOID) {
		return errors.New("error reading algorithm OID")
	}

	if !encryptionOID.Equal(OidRSAEncryption) {
		return errors.New("algorithm OID is not rsaEncryption")
	}

	if algorithm.Empty() {
		return errors.New("RSA algorithm identifier missing required NULL parameter")
	}

	var nullValue cryptobyte.String
	if !algorithm.ReadASN1(&nullValue, cryptobyte_asn1.NULL) {
		return errors.New("RSA algorithm identifier with non-NULL parameter")
	}

	if len(nullValue) != 0 {
		return errors.New("RSA algorithm identifier with NULL parameter containing data")
	}

	// ensure algorithm is empty and no trailing data is present
	if !algorithm.Empty() {
		return errors.New("RSA algorithm identifier with trailing data")
	}

	return errors.New("RSA algorithm appears correct, but didn't match byte-wise comparison")
}
