package filters

import "go/ast"

// Declarations takes in a list of a declarations and a predicate that takes in one declaration
// and returns a boolean. Only the declarations for which `predicate` returns true will be included in
// the returned list of declarations.
//
// For example, the following returns a list of only function declarations.
//
//	filters.Declarations(tree.Decls, func(decl ast.Decl) bool {
//		_, ok := decl.(*ast.FuncDecl)
//		return ok
//	})
//
// The order of declarations is maintained.
func Declarations(decls []ast.Decl, predicate func(decl ast.Decl) bool) (filtered []ast.Decl) {
	for _, decl := range decls {
		if predicate(decl) {
			filtered = append(filtered, decl)
		}
	}
	return
}

// FunctionsOnly returns a list of only the most outer function declarations present within
// the provided list. This filter does NOT recurse into those function declarations to find lambdas.
// For example, the following file...
//
//	func hi() bool {
//		return func() bool {
//			return true
//		}()
//	}
//
//	func hello() bool {
//		return false
//	}
//
// ...will return the hi and hello functions but not the inner lambda within hi.
func FunctionsOnly(decls []ast.Decl) []ast.Decl {
	return Declarations(decls, func(decl ast.Decl) bool {
		_, ok := decl.(*ast.FuncDecl)
		return ok
	})
}
