package rfc

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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
		Name            string
		Filename        string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "TestBeginsWithPeriodFQDN",
			Filename:       "beginsWithPeriodConstraintFQDN.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "TestIpAddressNotFQDN",
			Filename:        "ipAddressConstraintNotFQDN.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "certificate contained an inclusion name constraint that is not a fully qualified domain name: dns://192.168.1.1/ftp.example.org?type=A",
		},
		{
			Name:           "TestOnlyHostFQDN",
			Filename:       "onlyHostConstraintFQDN.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "TestNoAuthorityNotFQDN",
			Filename:        "noAuthorityConstraintNotFQDN.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "certificate contained an inclusion name constraint that is not a fully qualified domain name: example",
		},
		// Tests for the error messages
		{
			Name:            "Test1Exc1PermConstraint",
			Filename:        "exc1Perm1UriConstraints.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "certificate contained an inclusion name constraint that is not a fully qualified domain name: wrongHostConstraintExample2; certificate contained an exclusion name constraint that is not a fully qualified domain name: wrongHostConstraintExample",
		},
		{
			Name:            "TestMultExcMultPermConstraint",
			Filename:        "multExcMultPermUriConstraints.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "certificate contained multiple inclusion name constraints that are not fully qualified domain names: example3; example4; certificate contained multiple exclusion name constraints that are not fully qualified domain names: example; example2",
		},
		{
			Name:            "Test1ExcConstraint",
			Filename:        "exc1UriConstraint.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "certificate contained an exclusion name constraint that is not a fully qualified domain name: wrongHostConstraintExample",
		},
		{
			Name:            "TestMultExc1PermConstraints",
			Filename:        "multExc1PermUriConstraints.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "certificate contained an inclusion name constraint that is not a fully qualified domain name: example; certificate contained multiple exclusion name constraints that are not fully qualified domain names: wrongHost; example; wrongHost2",
		},
		{
			Name:            "TestMultPermConstraint",
			Filename:        "multPermUriConstraints.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "certificate contained multiple inclusion name constraints that are not fully qualified domain names: example; second; example",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			result := test.TestLint("e_name_constraint_not_fqdn", tc.Filename)

			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %v was %v", tc.ExpectedDetails, result.Details)
			}

			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result '%v' was '%v'", tc.ExpectedResult, result.Status)
			}

		})
	}
}
