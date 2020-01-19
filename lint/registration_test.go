package lint

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

import "testing"

func TestAllLintsHaveNameDescriptionSource(t *testing.T) {
	for name, lint := range Lints {
		if lint.Name == "" {
			t.Errorf("lint %s has empty name", name)
		}
		if lint.Description == "" {
			t.Errorf("lint %s has empty description", name)
		}
		if lint.Citation == "" {
			t.Errorf("lint %s has empty citation", name)
		}
	}
}

func TestAllLintsHaveSource(t *testing.T) {
	for name, lint := range Lints {
		if lint.Source == UnknownLintSource {
			t.Errorf("lint %s has unknown source", name)
		}
	}
}
