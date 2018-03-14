package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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
)

func TestExplicitTextUtf8NFC(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeExpTextUtf8.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtf8NotNFC(t *testing.T) {
	inputPath := "../testlint/testCerts/explicitTextUtf8NotNFC.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextBMPNFC(t *testing.T) {
	inputPath := "../testlint/testCerts/explicitTextBMPNFC.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextBMPNotNFC(t *testing.T) {
	inputPath := "../testlint/testCerts/explicitTextBMPNotNFC.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
