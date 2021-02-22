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

const want = `b7225d4ff9798f43ca4a5130b6fa3f100dd765cf24e02c9f802fabb0632bf302`

type NotCommitingGenTestCerts struct{}

func (i *NotCommitingGenTestCerts) CheckApplies(tree *ast.File, file *lint.File) bool {
	return strings.HasSuffix(file.Name, "genTestCerts.go")
}

func (i *NotCommitingGenTestCerts) Lint(tree *ast.File, file *lint.File) *lint.Result {
	contents, err := ioutil.ReadFile(file.Path)
	if err != nil {
		return lint.NewResult(fmt.Sprintf("failed to open %s", file.Name))
	}
	hasher := sha256.New()
	hasher.Write(contents)
	got := fmt.Sprintf("%x", hasher.Sum([]byte{}))
	if got == want {
		return nil
	}
	return lint.NewResult(fmt.Sprintf(`%s appears to have been modified and committed 
as a part of your change. This file is intended to be changed at your leisure, however we 
ask that these changed not be committed to the repo.

If you intended to submit changes to this file, then please run the following...

sha256sum v3/cmd/genTestCerts/genTestCerts.go

...and update the "want" constant in v3/integration/lints/lints/not_commiting_genTestCerts.go`, file.Path))
}
