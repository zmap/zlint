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
	"go/ast"
	"go/token"

	"github.com/zmap/zlint/v3/integration/lints/filters"
	"github.com/zmap/zlint/v3/integration/lints/lint"
)

type InitFirst struct{}

func (i *InitFirst) CheckApplies(tree *ast.File, file *lint.File) bool {
	return filters.IsALint(file)
}

func (i *InitFirst) Lint(tree *ast.File, file *lint.File) *lint.Result {
	functions := filters.FunctionsOnly(tree.Decls)
	if len(functions) == 0 {
		return lint.NewResult("Lint does not contain any functions or methods").
			AddCodeCitation(token.NoPos, token.NoPos, file)
	}
	// filters.FunctionsOnly have given us some guarantee that this type cast will succeed.
	firstFunction := functions[0].(*ast.FuncDecl)
	if isInit(firstFunction) {
		return nil
	}
	return lint.NewResult("Got the wrong method signature as the first function declaration within the linter.\n"+
		"ZLint lints must have func init() { ... } as their first function declaration").
		AddCodeCitation(firstFunction.Pos(), firstFunction.End(), file).
		SetCitations(
			"https://github.com/zmap/zlint/issues/371",
			"https://golang.org/doc/effective_go.html#init",
		)
}

func isInit(function *ast.FuncDecl) bool {
	isNamedInit := function.Name.Name == "init"
	isNotAMethod := function.Recv == nil
	hasNoParameters := len(function.Type.Params.List) == 0
	hasNoReturns := function.Type.Results == nil
	return isNamedInit && isNotAMethod && hasNoParameters && hasNoReturns
}
