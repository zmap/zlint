package rfc

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestInhibitAnyPolicyNotCrit(t *testing.T) {
	inputPath := "inhibitAnyNotCrit.pem"
	expected := lint.Error
	out := test.TestLint("e_inhibit_any_policy_not_critical", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestInhibitAnyPolicyCrit(t *testing.T) {
	inputPath := "inhibitAnyCrit.pem"
	expected := lint.Pass
	out := test.TestLint("e_inhibit_any_policy_not_critical", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
