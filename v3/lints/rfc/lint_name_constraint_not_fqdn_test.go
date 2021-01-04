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

func TestIsFQDN(t *testing.T) {
	inputPath := "constraintFQDN.pem"
	expected := lint.Pass
	out := test.TestLint("e_name_constraint_not_fqdn", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestBeginsWithPeridFQDN(t *testing.T) {
	inputPath := "beginsWithPeriodConstraintFQDN.pem"
	expected := lint.Error
	out := test.TestLint("e_name_constraint_not_fqdn", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIpAddressNotFQDN(t *testing.T) {
	inputPath := "ipAddressConstraintNotFQDN.pem"
	expected := lint.Error
	out := test.TestLint("e_name_constraint_not_fqdn", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestOnlyHostFQDN(t *testing.T) {
	inputPath := "onlyHostConstraintFQDN.pem"
	expected := lint.Error
	out := test.TestLint("e_name_constraint_not_fqdn", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNoAuthorityNotFQDN(t *testing.T) {
	inputPath := "noAuthorityConstraintNotFQDN.pem"
	expected := lint.Error
	out := test.TestLint("e_name_constraint_not_fqdn", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExcludedNameConstraintNotFQDN(t *testing.T) {
	inputPath := "excConstraintNotFQDN.pem"
	expected := lint.Error
	out := test.TestLint("e_name_constraint_not_fqdn", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
