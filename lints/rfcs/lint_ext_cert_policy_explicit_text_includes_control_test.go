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

func TestExplicitTextUtfControlX10(t *testing.T) {
	inputPath := "../testlint/testCerts/utf8ControlX10.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtfControlX88(t *testing.T) {
	inputPath := "../testlint/testCerts/utf8ControlX88.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtfNoControl(t *testing.T) {
	inputPath := "../testlint/testCerts/utf8NoControl.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
