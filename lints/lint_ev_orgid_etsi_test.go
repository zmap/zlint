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

func TestEvAltRegNumOrgidEtsi(t *testing.T) {
	m := map[string]LintStatus{
		"evAllGood.pem":                             NA,
		"EvAltRegNumCert54OrgIdInvalid.pem":         Error,
		"oiLEI.pem":                                 Pass,
		"EvAltRegNumCert56JurContryNotMatching.pem": Pass,
		"EvAltRegNumCert57NtrJurSopMissing.pem":     Pass,
		"EvAltRegNumCert58ValidNtr.pem":             Pass,
		"EvAltRegNumCert59Valid.pem":                Pass,
		"EvAltRegNumCert60OrgIdInvalid.pem":         Error,
		"EvAltRegNumCert61Valid.pem":                Pass,
		"EvAltRegNumCert62OrgIdLenZero.pem":         Error,
		"EvAltRegNumCert63Valid.pem":                Pass,
	}
	for inputPath, expected := range m {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["e_ev_orgid_etsi"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
