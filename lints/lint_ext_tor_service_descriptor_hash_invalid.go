/*
 * ZLint Copyright 2019 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package lints

import (
	"encoding/asn1"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
)

type torServiceDescHashInvalid struct{}

func (l *torServiceDescHashInvalid) Initialize() error {
	return nil
}

// CheckApplies returns true if the certificate is a subscriber certificate that
// contains a subject name ending in `.onion`.
func (l *torServiceDescHashInvalid) CheckApplies(c *x509.Certificate) bool {
	if !util.IsSubscriberCert(c) {
		return false
	}
	names := append(c.DNSNames, c.Subject.CommonName)
	for _, name := range names {
		if strings.HasSuffix(name, onionTLD) {
			return true
		}
	}
	return false
}

/*
 * The CAB Forum has created an extension of the TBSCertificate for use in
 * conveying hashes of keys related to .onion addresses.  The Tor Service
 * Descriptor Hash extension has the following format:
 *
 * cabf-TorServiceDescriptor OBJECT IDENTIFIER ::= { 2.23.140.1.31 }
 *
 * TorServiceDescriptorSyntax ::=
 * SEQUENCE ( 1..MAX ) of TorServiceDescriptorHash
 *
 * TorServiceDescriptorHash:: = SEQUENCE {
 *   onionURI              UTF8String
 *   algorithm             AlgorithmIdentifier
 *   subjectPublicKeyHash  BIT STRING
 * }
 *
 * Where the AlgorithmIdentifier is a hashing algorithm (defined in RFC 6234)
 * performed over the DER-encoding of an ASN.1 SubjectPublicKey of the .onion
 * service and SubjectPublicKeyHash is the hash output.
 */

// TorServiceDescriptorHash holds an onion URI (e.g. `http://<whatever>.onion`),
// a pkix.AlgorithmIdentifier, a slice of bytes holding a hash of the onion
// site's public key created with the specified hash algorithm, and the number
// of hash digest bits in the slice.
type TorServiceDescriptorHash struct {
	Onion    string
	Alg      pkix.AlgorithmIdentifier
	Hash     []byte
	HashBits int
}

// Valid returns nil if the the TorServiceDescriptorHash specifies a valid hash
// algorithm OID and the correct corresponding HashBits. Otherwise an error is
// returned describing the problem.
func (t TorServiceDescriptorHash) Valid() error {
	// If the length is <= 0 its definitely not valid.
	if t.HashBits <= 0 {
		return errors.New("invalid TorServiceDescriptorHash subjectPublicKeyHash, " +
			"bit length is <= 0")
	}

	// Check that the specified algorithm OID is an allowed hash algorithm OID.
	allowedHashAlgs := []asn1.ObjectIdentifier{
		util.SHA256OID,
		util.SHA384OID,
		util.SHA512OID,
	}
	goodHashAlg := util.SliceContainsOID(allowedHashAlgs, t.Alg.Algorithm)
	if !goodHashAlg {
		return fmt.Errorf("invalid TorServiceDescriptorHash algorithm %q",
			t.Alg.Algorithm)
	}

	// Check that the number of bits in the hash field match the algorithm
	// specified.
	if t.Alg.Algorithm.Equal(util.SHA256OID) && t.HashBits != 256 {
		return fmt.Errorf("invalid TorServiceDescriptorHash subjectPublicKeyHash, "+
			"alg is SHA256 but bit length is %d not %d",
			t.HashBits, 256)
	} else if t.Alg.Algorithm.Equal(util.SHA384OID) && t.HashBits != 384 {
		return fmt.Errorf("invalid TorServiceDescriptorHash subjectPublicKeyHash, "+
			"alg is SHA384 but bit length is %d not %d",
			t.HashBits, 384)
	} else if t.Alg.Algorithm.Equal(util.SHA512OID) && t.HashBits != 512 {
		return fmt.Errorf("invalid TorServiceDescriptorHash subjectPublicKeyHash, "+
			"alg is SHA512 but bit length is %d not %d",
			t.HashBits, 512)
	}
	return nil
}

// parseTorServiceDescriptorHash unmarshals a SEQUENCE from the provided data
// and parses a TorServiceDescriptorHash using the data contained in the
// sequence. The TorServiceDescriptorHash object and the remaining data are
// returned if no error occurs.
func parseTorServiceDescriptorHash(data []byte) (*TorServiceDescriptorHash, []byte, error) {
	// TorServiceDescriptorHash:: = SEQUENCE {
	//   onionURI UTF8String
	//   algorithm AlgorithmIdentifier
	//   subjectPublicKeyHash BIT STRING
	// }
	var outerSeq asn1.RawValue
	var err error
	data, err = asn1.Unmarshal(data, &outerSeq)
	if err != nil {
		return nil,
			data,
			errors.New("error unmarshaling TorServiceDescriptorHash SEQUENCE")
	}
	if outerSeq.Tag != asn1.TagSequence ||
		outerSeq.Class != asn1.ClassUniversal ||
		!outerSeq.IsCompound {
		return nil,
			data,
			errors.New("TorServiceDescriptorHash missing compound SEQUENCE tag")
	}
	fieldData := outerSeq.Bytes

	// Unmarshal and verify the structure of the onionURI UTF8String field.
	var rawOnionURI asn1.RawValue
	fieldData, err = asn1.Unmarshal(fieldData, &rawOnionURI)
	if err != nil {
		return nil,
			data,
			errors.New("error unmarshaling TorServiceDescriptorHash onionURI")
	}
	if rawOnionURI.Tag != asn1.TagUTF8String ||
		rawOnionURI.Class != asn1.ClassUniversal ||
		rawOnionURI.IsCompound {
		return nil,
			data,
			errors.New("TorServiceDescriptorHash missing non-compound UTF8String tag")
	}
	if !utf8.Valid(rawOnionURI.Bytes) {
		return nil,
			data,
			errors.New("TorServiceDescriptorHash UTF8String value was not valid UTF-8")
	}

	// Unmarshal and verify the structure of the algorithm UTF8String field.
	var algorithm pkix.AlgorithmIdentifier
	fieldData, err = asn1.Unmarshal(fieldData, &algorithm)
	if err != nil {
		return nil, nil, errors.New("error unmarshaling TorServiceDescriptorHash algorithm")
	}

	// Unmarshal and verify the structure of the Subject Public Key Hash BitString
	// field.
	var spkh asn1.BitString
	fieldData, err = asn1.Unmarshal(fieldData, &spkh)
	if err != nil {
		return nil, data, errors.New("error unmarshaling TorServiceDescriptorHash Hash")
	}

	// There should be no trailing data after the TorServiceDescriptorHash
	// SEQUENCE.
	if len(fieldData) > 0 {
		return nil, data, errors.New("trailing data after TorServiceDescriptorHash")
	}

	return &TorServiceDescriptorHash{
		Onion:    string(rawOnionURI.Bytes),
		Alg:      algorithm,
		HashBits: spkh.BitLength,
	}, data, nil
}

// parseTorServiceDescriptorSyntax parses the given pkix.Extension (assumed to
// have OID == util.BRTorServiceDescriptor) and returns a map of onion URIs to
// TorServiceDescriptorHash objects, or an error. An error will be returned if
// there are any structural errors related to the ASN.1 content (wrong tags,
// trailing data, missing fields, etc).
func parseTorServiceDescriptorSyntax(ext *pkix.Extension) (map[string]*TorServiceDescriptorHash, error) {
	baseErr := fmt.Sprintf(
		"certificate contained an invalid TorServiceDescriptor extension (oid %s)",
		util.BRTorServiceDescriptor.String())

	// TorServiceDescriptorSyntax ::=
	//    SEQUENCE ( 1..MAX ) of TorServiceDescriptorHash
	var seq asn1.RawValue
	rest, err := asn1.Unmarshal(ext.Value, &seq)
	if err != nil {
		return nil, fmt.Errorf("%s - unable to unmarshal outer SEQUENCE", baseErr)
	}
	if len(rest) != 0 {
		return nil, fmt.Errorf("%s - trailing data after outer SEQUENCE", baseErr)
	}
	if seq.Tag != asn1.TagSequence || seq.Class != asn1.ClassUniversal || !seq.IsCompound {
		return nil, fmt.Errorf("%s - invalid outer SEQUENCE", baseErr)
	}

	descriptors := make(map[string]*TorServiceDescriptorHash)
	rest = seq.Bytes
	for len(rest) > 0 {
		var descriptor *TorServiceDescriptorHash
		descriptor, rest, err = parseTorServiceDescriptorHash(rest)
		if err != nil {
			return nil, fmt.Errorf("%s - %s", baseErr, err)
		}
		descriptors[descriptor.Onion] = descriptor
	}
	return descriptors, nil
}

// Execute will lint the provided certificate. An Error LintResult will be
// returned if:
//   1) There is no TorServiceDescriptor extension present.
//   2) There are any problems parsing the TorServiceDescriptorSyntax or
//      TorServiceDescriptorHash values.
//   3) There are TorServiceDescriptorHash values specifing an invalid hash
//      algorithm or the wrong length of hash data.
//   4) There are any .onion domains without a corresponding
//      TorServiceDescriptorHash.
func (l *torServiceDescHashInvalid) Execute(c *x509.Certificate) *LintResult {
	ext := util.GetExtFromCert(c, util.BRTorServiceDescriptor)
	// If the BRTorServiceDescriptor extension is missing return a lint error. We
	// know the cert contains one or more `.onion` subjects because of
	// `CheckApplies` and all such certs are expected to have this extension.
	if ext == nil {
		return &LintResult{
			Status: Error,
			Details: fmt.Sprintf(
				"certificate contained a %s domain but is missing a TorServiceDescriptor "+
					"extension (oid %s)",
				onionTLD, util.BRTorServiceDescriptor.String()),
		}
	}
	// Parse the individual TorServiceDescriptorHash objects from the
	// TorServiceDescriptorSyntax SEQUENCE.
	descriptors, err := parseTorServiceDescriptorSyntax(ext)
	if err != nil {
		return &LintResult{
			Status:  Error,
			Details: err.Error(),
		}
	}
	// Validate the descriptors that were parsed
	for _, d := range descriptors {
		if err := d.Valid(); err != nil {
			return &LintResult{
				Status: Error,
				Details: fmt.Sprintf(
					"certificate contained an invalid TorServiceDescriptor extension (oid %s): %s",
					util.BRTorServiceDescriptor.String(),
					err.Error()),
			}
		}
	}
	var onionSubjectCount int
	for _, name := range append(c.DNSNames, c.Subject.CommonName) {
		if !strings.HasSuffix(name, onionTLD) {
			continue
		}
		onionSubjectCount++
		if _, found := descriptors["https://"+name]; !found {
			return &LintResult{
				Status: Error,
				Details: fmt.Sprintf(
					"%s domain name %q does not have a corresponding TorServiceDescriptorHash",
					onionTLD, name),
			}
		}
	}
	if onionSubjectCount != len(descriptors) {
		return &LintResult{
			Status: Error,
			Details: fmt.Sprintf(
				"certificate contained more TorServiceDescriptorHash entries than "+
					"%s domain names (%d vs %d)",
				onionTLD, len(descriptors), onionSubjectCount),
		}
	}
	return &LintResult{
		Status: Pass,
	}
}

func init() {
	RegisterLint(&Lint{
		Name:        "ext_tor_service_descriptor_hash_invalid",
		Description: "certificates with .onion names need valid TorServiceDescriptors in extension",
		// TODO(@cpu): Cite section of BRs instead of ballot?
		Citation:      "BRS: Ballot 144",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.OnionOnlyEVDate,
		Lint:          &torServiceDescHashInvalid{},
	})
}
