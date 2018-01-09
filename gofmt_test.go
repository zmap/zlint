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

package zlint

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestGofmt(t *testing.T) {
	globs := []string{
		"*.go",
		"cmd/*.go",
		"lints/*.go",
		"util/*.go",
	}
	for _, glob := range globs {
		gofmtCmd := "gofmt -s -l " + glob
		cmd := exec.Command("/bin/sh", "-c", gofmtCmd)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		if out.String() != "" {
			t.Errorf("glob %s not gofmt'ed", glob)
		}
	}
}
