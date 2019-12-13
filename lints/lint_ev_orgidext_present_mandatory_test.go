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
	"time"

	"github.com/zmap/zlint/util"
)

func TestEvAltRegNumOrgIdExtPresentMandatory(t *testing.T) {
	m := map[string]LintStatus{
		"EvAltRegNumCert52NoOrgId.pem": NA,
	}
	mBeforeExtMandDate := map[string]LintStatus{
		"EvAltRegNumCert53OrgIdInvalid.pem":         NE,
		"EvAltRegNumCert54OrgIdInvalid.pem":         NE,
		"EvAltRegNumCert55OrgIdExtMissing.pem":      NE,
		"EvAltRegNumCert56JurContryNotMatching.pem": NE,
	}
	mAfterExtMandDate := map[string]LintStatus{
		"EvAltRegNumCert55OrgIdExtMissing.pem":      Error,
		"EvAltRegNumCert67ValidNtrWithOrgIdExt.pem": Pass,
	}
	for inputPath, expected := range m {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["e_ev_orgidext_present_mandatory"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
	now := time.Now()
	var mCond map[string]LintStatus
	if now.Before(util.CABAltRegNumEvExtMandDate) {
		mCond = mBeforeExtMandDate
	} else {
		mCond = mAfterExtMandDate
	}
	for inputPath, expected := range mCond {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["e_ev_orgidext_present_mandatory"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
