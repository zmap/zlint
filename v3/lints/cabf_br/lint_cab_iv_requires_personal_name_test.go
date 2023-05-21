package cabf_br

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

func TestCertPolicyIvHasPerson(t *testing.T) {
	inputPath := "indivValGoodAllFields.pem"
	expected := lint.Pass
	out := test.TestLint("e_cab_iv_requires_personal_name", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyIvHasSurname(t *testing.T) {
	inputPath := "indivValSurnameOnly.pem"
	expected := lint.Error
	out := test.TestLint("e_cab_iv_requires_personal_name", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyIvHasLastName(t *testing.T) {
	inputPath := "indivValGivenNameOnly.pem"
	expected := lint.Error
	out := test.TestLint("e_cab_iv_requires_personal_name", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyIvNoPerson(t *testing.T) {
	inputPath := "indivValNoOrgOrPersonalNames.pem"
	expected := lint.Error
	out := test.TestLint("e_cab_iv_requires_personal_name", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
