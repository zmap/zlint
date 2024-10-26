package lints

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
	"crypto/sha256"
	"encoding/hex"
	"go/ast"
	"os"
	"strings"

	"github.com/zmap/zlint/v3/integration/lints/lint"
)

const want = `e113c11b7c4897c7e96579f175016094e48951a117b63c967d053e5ce83ec1cd`

type NotCommittingGenTestCerts struct{}

func (i *NotCommittingGenTestCerts) CheckApplies(tree *ast.File, file *lint.File) bool {
	return strings.HasSuffix(file.Name, "genTestCerts.go")
}

func (i *NotCommittingGenTestCerts) Lint(tree *ast.File, file *lint.File) *lint.Result {
	contents, err := os.ReadFile(file.Path)
	if err != nil {
		return lint.NewResult("failed to open " + file.Name)
	}
	hasher := sha256.New()
	_, err = hasher.Write(contents)
	if err != nil {
		return lint.NewResult("failed to hash the contents of " + file.Name)
	}
	got := hex.EncodeToString(hasher.Sum([]byte{}))
	if got == want {
		return nil
	}
	return lint.NewResult(file.Path + ` appears to have been modified and committed 
as a part of your change. This file is intended to be changed at your leisure, however we 
ask that these changed not be committed to the repo.

If you intended to submit changes to this file, then please run the following...

sha256sum cmd/genTestCerts/genTestCerts.go

...and update the "want" constant in v3/integration/lints/lints/not_committing_genTestCerts.go`)
}
