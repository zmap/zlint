package util

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

import (
	"testing"

	"github.com/zmap/zcrypto/encoding/asn1"
)

// test-only mirrors of the PSD2 ASN.1 types. testPSD2QcTypeMalformed
// deliberately tags NCAName as PrintableString instead of UTF8String to
// produce a statement that parses but does not round-trip to the same
// bytes the real UTF8String-tagged PSD2QcType would produce.
type testRoleOfPSP struct {
	RoleOfPspOid  asn1.ObjectIdentifier
	RoleOfPspName string `asn1:"utf8"`
}

type testPSD2QcTypeValid struct {
	RolesOfPSP []testRoleOfPSP
	NCAName    string `asn1:"utf8"`
	NCAId      string `asn1:"utf8"`
}

type testPSD2QcTypeMalformed struct {
	RolesOfPSP []testRoleOfPSP
	NCAName    string `asn1:"printable"`
	NCAId      string `asn1:"utf8"`
}

type testQcStatement struct {
	Oid  asn1.ObjectIdentifier
	Info asn1.RawValue
}

// buildPsd2ExtValue wraps already-encoded PSD2QcType bytes into a
// one-statement QcStatements extension value, i.e. what CheckApplies /
// ParseQcStatem expect to receive as extVal.
func buildPsd2ExtValue(t *testing.T, psd2Bytes []byte) []byte {
	t.Helper()
	stmt := testQcStatement{
		Oid:  IdEtsiPsd2Statem,
		Info: asn1.RawValue{FullBytes: psd2Bytes},
	}
	extVal, err := asn1.Marshal([]testQcStatement{stmt})
	if err != nil {
		t.Fatalf("failed to marshal QcStatements extension value: %v", err)
	}
	return extVal
}

func TestParseQcStatemPsd2Valid(t *testing.T) {
	psd2 := testPSD2QcTypeValid{
		RolesOfPSP: []testRoleOfPSP{
			{RoleOfPspOid: asn1.ObjectIdentifier{0, 4, 0, 19495, 1, 1}, RoleOfPspName: "PSP_AS"},
		},
		NCAName: "Banco de España",
		NCAId:   "ES-BDE",
	}
	psd2Bytes, err := asn1.Marshal(psd2)
	if err != nil {
		t.Fatalf("failed to marshal PSD2QcType: %v", err)
	}
	extVal := buildPsd2ExtValue(t, psd2Bytes)

	result := ParseQcStatem(extVal, IdEtsiPsd2Statem)
	if !result.IsPresent() {
		t.Fatalf("expected PSD2 QC statement to be present")
	}
	if result.GetErrorInfo() != "" {
		t.Fatalf("expected no error info for valid PSD2 QC statement, got: %q", result.GetErrorInfo())
	}
	psd2Result, ok := result.(EtsiPsd2)
	if !ok {
		t.Fatalf("expected result to be of type EtsiPsd2, got %T", result)
	}
	if len(psd2Result.Decoded.RolesOfPSP) != 1 {
		t.Fatalf("expected 1 role, got %d", len(psd2Result.Decoded.RolesOfPSP))
	}
	if psd2Result.Decoded.NCAId != "ES-BDE" {
		t.Fatalf("expected NCAId %q, got %q", "ES-BDE", psd2Result.Decoded.NCAId)
	}
}

func TestParseQcStatemPsd2MalformedEncoding(t *testing.T) {
	psd2 := testPSD2QcTypeMalformed{
		RolesOfPSP: []testRoleOfPSP{
			{RoleOfPspOid: asn1.ObjectIdentifier{0, 4, 0, 19495, 1, 1}, RoleOfPspName: "PSP_AS"},
		},
		NCAName: "Banco de Espana",
		NCAId:   "ES-BDE",
	}
	psd2Bytes, err := asn1.Marshal(psd2)
	if err != nil {
		t.Fatalf("failed to marshal malformed PSD2QcType: %v", err)
	}
	extVal := buildPsd2ExtValue(t, psd2Bytes)

	result := ParseQcStatem(extVal, IdEtsiPsd2Statem)
	if result.GetErrorInfo() == "" {
		t.Fatalf("expected error info for PSD2 QC statement with wrong string type encoding, got none")
	}
}

func TestParseQcStatemPsd2UnmarshalFailure(t *testing.T) {
	psd2 := testPSD2QcTypeValid{
		RolesOfPSP: []testRoleOfPSP{
			{RoleOfPspOid: asn1.ObjectIdentifier{0, 4, 0, 19495, 1, 1}, RoleOfPspName: "PSP_AS"},
		},
		NCAName: "Banco de España",
		NCAId:   "ES-BDE",
	}
	psd2Bytes, err := asn1.Marshal(psd2)
	if err != nil {
		t.Fatalf("failed to marshal PSD2QcType: %v", err)
	}
	// Truncate the otherwise-valid, already-marshaled PSD2QcType bytes by one
	// byte so the outer SEQUENCE's declared length no longer matches the
	// available content. Note: appending extra trailing bytes instead does
	// NOT work here, because zcrypto's asn1 fork tolerates unconsumed
	// trailing content within a declared SEQUENCE length, and any bytes
	// appended past the declared length get silently dropped when the
	// RawValue capturing statem.Any.FullBytes reads exactly one TLV in
	// buildPsd2ExtValue's wrapping. Truncation instead makes the declared
	// length exceed the available bytes, which reliably makes
	// asn1.Unmarshal fail with "data truncated" -- exercising the
	// "error parsing the statementInfo field" branch of ParseQcStatem
	// (asn1.Unmarshal(statem.Any.FullBytes, &etsiObj.Decoded) failing
	// outright), as opposed to the checkAsn1Reencoding round-trip branch
	// already covered by TestParseQcStatemPsd2MalformedEncoding.
	truncated := psd2Bytes[:len(psd2Bytes)-1]
	extVal := buildPsd2ExtValue(t, truncated)

	result := ParseQcStatem(extVal, IdEtsiPsd2Statem)
	if result.GetErrorInfo() == "" {
		t.Fatalf("expected error info for PSD2 QC statement with truncated statementInfo bytes, got none")
	}
}

func TestParseQcStatemPsd2NotPresent(t *testing.T) {
	extVal := buildPsd2ExtValue(t, mustMarshal(t, testPSD2QcTypeValid{
		RolesOfPSP: []testRoleOfPSP{{RoleOfPspOid: asn1.ObjectIdentifier{0, 4, 0, 19495, 1, 1}, RoleOfPspName: "PSP_AS"}},
		NCAName:    "Banco de España",
		NCAId:      "ES-BDE",
	}))
	// Ask for a different (unrelated, but already-registered) statement OID:
	// the PSD2 statement is present in extVal but we're not asking about it,
	// so IsPresent() must be false.
	result := ParseQcStatem(extVal, IdEtsiQcsQcCompliance)
	if result.IsPresent() {
		t.Fatalf("expected IdEtsiQcsQcCompliance to be absent from an extension containing only a PSD2 statement")
	}
}

func mustMarshal(t *testing.T, v interface{}) []byte {
	t.Helper()
	b, err := asn1.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}
	return b
}
