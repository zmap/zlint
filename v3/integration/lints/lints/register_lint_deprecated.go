package lints

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

import (
	"go/ast"

	"github.com/zmap/zlint/v3/integration/lints/filters"
	"github.com/zmap/zlint/v3/integration/lints/lint"
)

type RegisterLintDeprecated struct{}

func (r *RegisterLintDeprecated) CheckApplies(tree *ast.File, file *lint.File) bool {
	return filters.IsALint(file)
}

func (r *RegisterLintDeprecated) Lint(tree *ast.File, file *lint.File) *lint.Result {
	var result *lint.Result
	visitor := &selectorExprVisitor{fn: func(expr *ast.SelectorExpr, node ast.Node) {
		if expr.Sel.Name != "RegisterLint" {
			return
		}
		result = lint.NewResult("lint.RegisterLint is deprecated and should not be used. "+
			"Please use the register function specific to your lint classification (I.E. "+
			"lint.RegisterCertificateLint for certificate lints and lint.RegisterRevocationListLint for CRL lints).").
			AddCodeCitation(node.Pos(), node.End(), file).
			SetCitations("https://github.com/zmap/zlint/issues/765")
	}}
	ast.Walk(visitor, tree)
	return result
}

type selectorExprVisitor struct {
	fn func(expr *ast.SelectorExpr, node ast.Node)
}

func (v *selectorExprVisitor) Visit(node ast.Node) ast.Visitor {
	selectorExpr, ok := node.(*ast.SelectorExpr)
	if !ok {
		return v
	}
	v.fn(selectorExpr, node)
	return nil
}
