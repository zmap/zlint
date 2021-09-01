package lints

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
	"crypto/sha256"
	"fmt"
	"go/ast"
	"io/ioutil"
	"strings"

	"github.com/zmap/zlint/v3/integration/lints/lint"
)

const want = `25f02402841d94ab154bf3e8856dbf83fa98ef8f27910d09772b72da221ad882`

type NotCommittingGenTestCerts struct{}

func (i *NotCommittingGenTestCerts) CheckApplies(tree *ast.File, file *lint.File) bool {
	return strings.HasSuffix(file.Name, "genTestCerts.go")
}

func (i *NotCommittingGenTestCerts) Lint(tree *ast.File, file *lint.File) *lint.Result {
	contents, err := ioutil.ReadFile(file.Path)
	if err != nil {
		return lint.NewResult(fmt.Sprintf("failed to open %s", file.Name))
	}
	hasher := sha256.New()
	_, err = hasher.Write(contents)
	if err != nil {
		return lint.NewResult(fmt.Sprintf("failed to hash the contents of %s", file.Name))
	}
	got := fmt.Sprintf("%x", hasher.Sum([]byte{}))
	if got == want {
		return nil
	}
	return lint.NewResult(fmt.Sprintf(`%s appears to have been modified and committed 
as a part of your change. This file is intended to be changed at your leisure, however we 
ask that these changed not be committed to the repo.

If you intended to submit changes to this file, then please run the following...

sha256sum cmd/genTestCerts/genTestCerts.go

...and update the "want" constant in v3/integration/lints/lints/not_committing_genTestCerts.go`, file.Path))
}
