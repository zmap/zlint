package rfc

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestExplicitTextUtf8NFC(t *testing.T) {
	inputPath := "userNoticeExpTextUtf8.pem"
	expected := lint.Pass
	out := test.TestLint("w_ext_cert_policy_explicit_text_not_nfc", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtf8NotNFC(t *testing.T) {
	inputPath := "explicitTextUtf8NotNFC.pem"
	expected := lint.Warn
	out := test.TestLint("w_ext_cert_policy_explicit_text_not_nfc", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextBMPNFC(t *testing.T) {
	inputPath := "explicitTextBMPNFC.pem"
	expected := lint.Pass
	out := test.TestLint("w_ext_cert_policy_explicit_text_not_nfc", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextBMPNotNFC(t *testing.T) {
	inputPath := "explicitTextBMPNotNFC.pem"
	expected := lint.Warn
	out := test.TestLint("w_ext_cert_policy_explicit_text_not_nfc", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
