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

func TestUriNameConstraintsFqdn(t *testing.T) {
	testCases := []struct {
		Name           string
		Filename       string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "TestBeginsWithPeridFQDN",
			Filename:       "beginsWithPeriodConstraintFQDN.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "TestIpAddressNotFQDN",
			Filename:       "ipAddressConstraintNotFQDN.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TestOnlyHostFQDN",
			Filename:       "onlyHostConstraintFQDN.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "TestNoAuthorityNotFQDN",
			Filename:       "noAuthorityConstraintNotFQDN.pem",
			ExpectedResult: lint.Error,
		},
		// Tests for the error messages
		{
			Name:           "Test1Exc1PermConstraint",
			Filename:       "exc1Perm1UriConstraints.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TestMultExcMultPermConstraint",
			Filename:       "multExcMultPermUriConstraints.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Test1ExcConstraint",
			Filename:       "exc1UriConstraint.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TestMultExc1PermConstraints",
			Filename:       "multExc1PermUriConstraints.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TestMultPermConstraint",
			Filename:       "multPermUriConstraints.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			result := test.TestLint("e_name_constraint_not_fqdn", tc.Filename)
			if tc.ExpectedResult == lint.Error {
				t.Logf("error message for %v: %v", tc.Name, result.Details)
			}

			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}

		})
	}
}
