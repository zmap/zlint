/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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

package rfc

import (
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
	"golang.org/x/crypto/ocsp"

	"encoding/asn1"
)

func init() {
	lint.RegisterOcspResponseLint(&lint.OcspResponseLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_ocsp_empty_seq_in_extns",
			Description:   "Checks that Extensions-type fields of OCSP responses are NOT an empty SEQUENCE",
			Citation:      "RFC 6960, Appendix B",
			Source:        lint.RFC6960,
			EffectiveDate: util.RFC6960Date,
		},
		Lint: NewOCSPEmptySequenceInExtensions,
	})
}

type ocspResponse struct {
	ResponseStatus asn1.Enumerated
	ResponseBytes  responseBytesType `asn1:"tag:0,explicit,optional"`
}

type responseBytesType struct {
	ResponseType asn1.ObjectIdentifier
	Response     []byte
}

type basicOCSPResponse struct {
	TbsResponseData    responseData
	SignatureAlgorithm asn1.RawValue
	Signature          asn1.RawValue
	Certs              asn1.RawValue `asn1:"optional"`
}

type responseData struct {
	Version            int `asn1:"tag:0,explicit,optional"`
	ResponderID        asn1.RawValue
	ProducedAt         asn1.RawValue
	Responses          []singleResponse
	ResponseExtensions asn1.RawValue `asn1:"optional"`
}

type singleResponse struct {
	CertID           asn1.RawValue
	CertStatus       asn1.RawValue
	ThisUpdate       asn1.RawValue
	NextUpdate       asn1.RawValue `asn1:"optional"`
	SingleExtensions asn1.RawValue `asn1:"optional"`
}

type OCSPEmptySequenceInExtensions struct{}

func NewOCSPEmptySequenceInExtensions() lint.OcspResponseLintInterface {
	return &OCSPEmptySequenceInExtensions{}
}

func (l *OCSPEmptySequenceInExtensions) CheckApplies(c *ocsp.Response) bool {
	return true
}

// From RFC 6960
var idPkixOcspBasic = "1.3.6.1.5.5.7.48.1.1"
var OCSPResponseStatusSuccessful = asn1.Enumerated(0)

func (l *OCSPEmptySequenceInExtensions) Execute(c *ocsp.Response) *lint.LintResult {

	ocspResp := ocspResponse{}
	_, err := asn1.Unmarshal(c.Raw, &ocspResp)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal,
			Details: "Could not decode OCSP Response: " + err.Error()}
	}

	if ocspResp.ResponseStatus != OCSPResponseStatusSuccessful {
		return &lint.LintResult{Status: lint.NA}
	}

	respType := ocspResp.ResponseBytes.ResponseType
	if respType.String() != idPkixOcspBasic {
		return &lint.LintResult{Status: lint.Error,
			Details: "Unexpected OCSP Response type " + respType.String(),
		}
	}

	basicResp := basicOCSPResponse{}
	_, err = asn1.Unmarshal(ocspResp.ResponseBytes.Response, &basicResp)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal,
			Details: "Could not decode OCSP Response: " + err.Error()}
	}

	// Check extensions associated to the entire ResponseData
	r := basicResp.TbsResponseData
	if len(r.ResponseExtensions.Bytes) > 1 && // Field was found
		r.ResponseExtensions.Bytes[0] == 0x30 && // and it's a SEQUENCE
		r.ResponseExtensions.Bytes[1] == 0 { // with zero elements
		return &lint.LintResult{Status: lint.Error,
			Details: "Empty SEQUENCE in the ResponseExtensions field"}
	}

	// Check extensions associated to each SingleResponse
	for _, sr := range basicResp.TbsResponseData.Responses {
		if len(sr.SingleExtensions.Bytes) > 1 && // Field was found
			sr.SingleExtensions.Bytes[0] == 0x30 && // and it's a SEQUENCE
			sr.SingleExtensions.Bytes[1] == 0 { // with zero elements
			return &lint.LintResult{Status: lint.Error,
				Details: "Empty SEQUENCE in the SingleExtensions field"}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
