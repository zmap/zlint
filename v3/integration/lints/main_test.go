package main

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

import "testing"

// `main` has a runner function simply named `run` which takes in a string which is a directory that will be recursively
// searched for Go files to lint. In this particular case, we have some sample Go files under `maintestdata`.
func TestFullRun(t *testing.T) {
	results, err := run("testdata")
	if err != nil {
		t.Error(err)
		return
	}
	if len(results) != 1 {
		t.Errorf("expected 1 error, got %d", len(results))
	}
}
