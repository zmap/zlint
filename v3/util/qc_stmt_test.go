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
	"strings"
	"testing"

	"github.com/zmap/zcrypto/encoding/asn1"
)

// testPSD2QcTypeMalformed deliberately tags NCAName as PrintableString
// instead of UTF8String to produce a statement that parses but does not
// round-trip to the same bytes the real UTF8String-tagged PSD2QcType would
// produce. RoleOfPSP and PSD2QcType themselves are used directly from
// qc_stmt.go below; this file is package util, so it has direct access to
// the real types and doesn't need mirrors of them.
type testPSD2QcTypeMalformed struct {
	RolesOfPSP []RoleOfPSP
	NCAName    string `asn1:"printable"`
	NCAId      string `asn1:"utf8"`
}

// buildPsd2ExtValue wraps already-encoded PSD2QcType bytes into a
// one-statement QcStatements extension value, i.e. what CheckApplies /
// ParseQcStatem expect to receive as extVal.
func buildPsd2ExtValue(t *testing.T, psd2Bytes []byte) []byte {
	t.Helper()
	stmt := qcStatementWithInfoField{
		Oid: IdEtsiPsd2Statem,
		Any: asn1.RawValue{FullBytes: psd2Bytes},
	}
	return mustMarshal(t, []qcStatementWithInfoField{stmt})
}

// validPsd2QcType returns a well-formed PSD2QcType shared by every test in
// this file that needs one but isn't itself testing a specific field value.
func validPsd2QcType() PSD2QcType {
	return PSD2QcType{
		RolesOfPSP: []RoleOfPSP{
			{RoleOfPspOid: asn1.ObjectIdentifier{0, 4, 0, 19495, 1, 1}, RoleOfPspName: "PSP_AS"},
		},
		NCAName: "Banco de España",
		NCAId:   "ES-BDE",
	}
}

func TestParseQcStatemPsd2Valid(t *testing.T) {
	extVal := buildPsd2ExtValue(t, mustMarshal(t, validPsd2QcType()))

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
		RolesOfPSP: []RoleOfPSP{
			{RoleOfPspOid: asn1.ObjectIdentifier{0, 4, 0, 19495, 1, 1}, RoleOfPspName: "PSP_AS"},
		},
		NCAName: "Banco de Espana",
		NCAId:   "ES-BDE",
	}
	extVal := buildPsd2ExtValue(t, mustMarshal(t, psd2))

	result := ParseQcStatem(extVal, IdEtsiPsd2Statem)
	if result.GetErrorInfo() == "" {
		t.Fatalf("expected error info for PSD2 QC statement with wrong string type encoding, got none")
	}
}

func TestParseQcStatemPsd2UnmarshalFailure(t *testing.T) {
	psd2Bytes := mustMarshal(t, validPsd2QcType())
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

// TestParseQcStatemPsd2SizeConstraints exercises the Annex A SIZE(1..256)
// bound on NCAName, NCAId, and RoleOfPspName: exactly at the min/max bounds
// must pass, and one below/above must fail. Also confirms the bound is
// counted in Unicode characters, not bytes, since these are UTF8Strings.
func TestParseQcStatemPsd2SizeConstraints(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(psd2 *PSD2QcType)
		wantErr bool
	}{
		{"NCAName at min length (1) passes", func(psd2 *PSD2QcType) { psd2.NCAName = "A" }, false},
		{"NCAName at max length (256) passes", func(psd2 *PSD2QcType) { psd2.NCAName = strings.Repeat("A", 256) }, false},
		{"NCAName empty fails", func(psd2 *PSD2QcType) { psd2.NCAName = "" }, true},
		{"NCAName over max length (257) fails", func(psd2 *PSD2QcType) { psd2.NCAName = strings.Repeat("A", 257) }, true},
		{"NCAId empty fails", func(psd2 *PSD2QcType) { psd2.NCAId = "" }, true},
		{"NCAId over max length (257) fails", func(psd2 *PSD2QcType) { psd2.NCAId = strings.Repeat("A", 257) }, true},
		{"RoleOfPspName empty fails", func(psd2 *PSD2QcType) { psd2.RolesOfPSP[0].RoleOfPspName = "" }, true},
		{"RoleOfPspName over max length (257) fails", func(psd2 *PSD2QcType) { psd2.RolesOfPSP[0].RoleOfPspName = strings.Repeat("A", 257) }, true},
		{"256-character NCAName counted in runes, not bytes, passes", func(psd2 *PSD2QcType) { psd2.NCAName = strings.Repeat("é", 256) }, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			psd2 := validPsd2QcType()
			tt.mutate(&psd2)
			extVal := buildPsd2ExtValue(t, mustMarshal(t, psd2))

			result := ParseQcStatem(extVal, IdEtsiPsd2Statem)
			gotErr := result.GetErrorInfo() != ""
			if gotErr != tt.wantErr {
				t.Errorf("GetErrorInfo() = %q, wantErr %v", result.GetErrorInfo(), tt.wantErr)
			}
		})
	}
}

func TestParseQcStatemPsd2NotPresent(t *testing.T) {
	extVal := buildPsd2ExtValue(t, mustMarshal(t, validPsd2QcType()))
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
