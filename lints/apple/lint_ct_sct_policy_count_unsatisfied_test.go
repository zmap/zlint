/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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

package lints

import (
	"fmt"
	"testing"
)

func TestSCTCountPolicyUnsatisified(t *testing.T) {
	// NOTE(@cpu): Hello future human. If you need to recreate any of the
	// Filenames referenced in this test you will need the `sctTestCerts.go`
	// program[0]. Each test case has a comment that includes the invocation
	// arguments that were used to create the test file.
	//
	// [0]: https://gist.github.com/cpu/6d26b2718f29e184ff88a90f02d7cbcb
	testCases := []struct {
		Name           string
		Filename       string
		ExpectedResult LintStatus
	}{
		{
			Name: "No SCTs, poisoned",
			// go run sctTestCerts.go -lifetime 3 -scts 0 -poison > testlint/testCerts/ctNoSCTsPoisoned.pem
			Filename:       "ctNoSCTsPoisoned.pem",
			ExpectedResult: NA,
		},
		{
			Name: "No SCTs, no poison",
			// go run sctTestCerts.go -lifetime 3 -scts 0 > testlint/testCerts/ctNoSCTs.pem
			Filename:       "ctNoSCTs.pem",
			ExpectedResult: Notice,
		},
		{
			Name: "Lifetime <15mo, 1 SCT",
			// go run sctTestCerts.go -lifetime 3 -scts 1 > testlint/testCerts/ct3mo1SCTs.pem
			Filename:       "ct3mo1SCTs.pem",
			ExpectedResult: Notice,
		},
		{
			Name: "Lifetime <15mo, 2 SCTs diff logs",
			// go run sctTestCerts.go -lifetime 3 -scts 2 > testlint/testCerts/ct3mo2SCTs.pem
			Filename:       "ct3mo2SCTs.pem",
			ExpectedResult: Pass,
		},
		{
			Name: "Lifetime <15mo, 2 SCTs same logs",
			// go run sctTestCerts.go -lifetime 3 -scts 2 -differentLogs=false > testlint/testCerts/ct3mo2DupeSCTs.pem
			Filename:       "ct3mo2DupeSCTs.pem",
			ExpectedResult: Notice,
		},
		{
			Name: "Lifetime >15mo <27mo, 2 SCTs diff logs",
			// go run sctTestCerts.go -lifetime 18 -scts 2 > testlint/testCerts/ct18mo2SCTs.pem
			Filename:       "ct18mo2SCTs.pem",
			ExpectedResult: Notice,
		},
		{
			Name: "Lifetime >15mo <27mo, 3 SCTs diff logs",
			// go run sctTestCerts.go -lifetime 18 -scts 3 > testlint/testCerts/ct18mo3SCTs.pem
			Filename:       "ct18mo3SCTs.pem",
			ExpectedResult: Pass,
		},
		{
			Name: "Lifetime >27mo <39mo, 3 SCTs diff logs",
			// go run sctTestCerts.go -lifetime 38 -scts 3 > testlint/testCerts/ct38mo3SCTs.pem
			Filename:       "ct38mo3SCTs.pem",
			ExpectedResult: Notice,
		},
		{
			Name: "Lifetime >27mo <39mo, 4 SCTs diff logs",
			// go run sctTestCerts.go -lifetime 38 -scts 4 > testlint/testCerts/ct38mo4SCTs.pem
			Filename:       "ct38mo4SCTs.pem",
			ExpectedResult: Pass,
		},
		{
			Name: "Lifetime >39mo, 4 SCTs diff logs",
			// go run sctTestCerts.go -lifetime 666 -scts 4 > testlint/testCerts/ct666mo4SCTs.pem
			Filename:       "ct666mo4SCTs.pem",
			ExpectedResult: Notice,
		},
		{
			Name: "Lifetime >39mo, 5 SCTs diff logs",
			// go run sctTestCerts.go -lifetime 666 -scts 5 > testlint/testCerts/ct666mo5SCTs.pem
			Filename:       "ct666mo5SCTs.pem",
			ExpectedResult: Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.Filename)
			result := Lints["w_ct_sct_policy_count_unsatisfied"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
