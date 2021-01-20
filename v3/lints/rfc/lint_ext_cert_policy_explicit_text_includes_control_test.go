package rfc

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

func TestExplicitTextUtfControlX10(t *testing.T) {
	inputPath := "utf8ControlX10.pem"
	expected := lint.Warn
	out := test.TestLint("w_ext_cert_policy_explicit_text_includes_control", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtfControlX88(t *testing.T) {
	inputPath := "utf8ControlX88.pem"
	expected := lint.Warn
	out := test.TestLint("w_ext_cert_policy_explicit_text_includes_control", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtfNoControl(t *testing.T) {
	inputPath := "utf8NoControl.pem"
	expected := lint.Pass
	out := test.TestLint("w_ext_cert_policy_explicit_text_includes_control", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
