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

package profiles

import (
	"io/ioutil"
	"testing"

	"github.com/zmap/zlint/v3/lint"
	_ "github.com/zmap/zlint/v3/lints/apple"
	_ "github.com/zmap/zlint/v3/lints/cabf_br"
	_ "github.com/zmap/zlint/v3/lints/cabf_cs_br"
	_ "github.com/zmap/zlint/v3/lints/cabf_ev"
	_ "github.com/zmap/zlint/v3/lints/cabf_smime_br"
	_ "github.com/zmap/zlint/v3/lints/community"
	_ "github.com/zmap/zlint/v3/lints/etsi"
	_ "github.com/zmap/zlint/v3/lints/mozilla"
	_ "github.com/zmap/zlint/v3/lints/rfc"
)

// We would like to make sure that there is a generic test that makes sure
// that all profiles actually refer to registered lints.
func TestLintsInAllProfilesExist(t *testing.T) {
	for _, profile := range lint.AllProfiles() {
		for _, l := range profile.LintNames {
			if lint.GlobalRegistry().ByName(l) == nil {
				t.Errorf("Profile '%s' declares lint '%s' which does not exist", profile.Name, l)
			}
		}
	}
}

// In order to run TestLintsInAllProfilesExist we need to import all lint source packages in order
// to run their init functions. This test makes sure that if anyone adds a new
// lint source in the future that we don't miss importing it into this test file.
func TestNotMissingAnyLintSources(t *testing.T) {
	expected := map[string]bool{
		"apple":         true,
		"cabf_br":       true,
		"cabf_cs_br":    true,
		"cabf_ev":       true,
		"cabf_smime_br": true,
		"community":     true,
		"etsi":          true,
		"mozilla":       true,
		"rfc":           true,
	}
	dir, err := ioutil.ReadDir("../lints")
	if err != nil {
		t.Fatal(err)
	}
	for _, info := range dir {
		if !info.IsDir() {
			continue
		}
		if _, ok := expected[info.Name()]; !ok {
			t.Errorf("We need to import each lint source in order to ensure that all lint names referred to by "+
				"declared profiles actually exist. However, we found the directory lints/%s which is not a lint "+
				"source that this test is aware of. Please add the following import to the top if this test file: "+
				"_ \"github.com/zmap/zlint/v3/lints/%s\"", info.Name(), info.Name())
		}
	}

}
