package lint

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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type Lint interface {
	Lint(tree *ast.File, file *File) *Result
	CheckApplies(tree *ast.File, file *File) bool
}

// A Result encodes any unmet expectation laid out by your lint. It consists of a single message, a list of code
// citations, and a list of lint citations.
//
// The message should be succinct and descriptive of the core issue. This message can only be set in the constructor,
// NewResult. For example...
//
//		"Go style guides suggest not using bare returns in complex functions"
//
// Code citations are the locations within the file that did not meet your expectations. Please see AddCodeCitations
// for information on how to add these to the Result type. Adding a code citation will result in the file, line number
// and raw source code appearing in the lint result. For example...
//
//	File ../../lints/cabf_br/lint_cab_dv_conflicts_with_locality.go, line 28
//
//	func (l *certPolicyConflictsWithLocality) Initialize() error {
//		return nil
//	}
//
// The lint citations are additional information to help the contributor understand why their code failed
// this lint and, if possible, some hints or resources on how to correct the issue. Every citation will be listed on its
// own line.
type Result struct {
	message       string
	codeCitations []string
	citations     []string
}

func NewResult(message string) *Result {
	return &Result{message: message}
}

// AddCodeCitation takes the starting and ending position of a block of code within a file.
// Upon calling the String method, every code citation will be printed alongside the
// result. This code citation lists the file and line of the code in question
// as well as the raw block of source code.
//
// For example:
//
//	File ../../lints/cabf_br/lint_cab_dv_conflicts_with_locality.go, line 28
//
//	func (l *certPolicyConflictsWithLocality) Initialize() error {
//		return nil
//	}
//
//
func (r *Result) AddCodeCitation(start, end token.Pos, file *File) *Result {
	srcCode := make([]byte, end-start)
	reader := strings.NewReader(file.Src)
	// We have no real interest in the error return since this is an in-memory reader.
	_, _ = reader.ReadAt(srcCode, int64(start))
	lineno := file.LineOf(start)
	citation := fmt.Sprintf("File %s, line %d\n\n%s\n\n", file.Path, lineno, string(srcCode))
	r.codeCitations = append(r.codeCitations, citation)
	return r
}

// SetCitations sets a list of citations that users can reference in order to understand
// the error that they received. Upon calling the String method each citation will be
// listed on their on own line.
//
// For example:
//
//	For more information, please see the following citations.
//		https://github.com/zmap/zlint/issues/371
//		https://golang.org/doc/effective_go.html#init
//
// The above links a GitHub issue that discuss the lint in question as well as a link
// to Golang's magic `init` method (because the lint in question is asking the contributor
// to implement `init` at a particular spot in the file).
func (r *Result) SetCitations(citations ...string) *Result {
	r.citations = citations
	return r
}

func (r *Result) String() string {
	b := strings.Builder{}
	b.WriteString("--------------------\n")
	b.WriteString("Linting Error\n\n")
	b.WriteString(r.message)
	b.WriteString("\n\n")
	for _, code := range r.codeCitations {
		b.WriteString(code)
	}
	if len(r.citations) > 0 {
		b.WriteString("For more information, please see the following citations.\n")
	}
	for _, citation := range r.citations {
		b.WriteByte('\t')
		b.WriteString(citation)
		b.WriteByte('\n')
	}
	return b.String()
}

type File struct {
	Src   string
	Path  string
	Name  string
	Lines []string
}

// LineOf computes which line a particular position within a file lands on.
//
//	This is not the greatest song in the world.
// 	No, this is just a tribute.
//	Couldn't remember the greatest song in the world.
// 	No, this is just a tribute!
//
// The word "remember" begins at position 81 within this text, therefor LineOf(81) should return line 3.
func (f *File) LineOf(pos token.Pos) int {
	start := 0
	end := 0
	for lineno, line := range f.Lines {
		start = end
		end = start + len(line)
		if int(pos) >= start && int(pos) <= end {
			return lineno + 1
		}
	}
	return int(token.NoPos)
}

func NewFile(name, src string) *File {
	return &File{src, name, filepath.Base(name), strings.Split(src, "\n")}
}

func Parse(path string) (*ast.File, *File, error) {
	fset := new(token.FileSet)
	tree, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	file := NewFile(path, string(b))
	return tree, file, nil
}

func RunLintForFile(path string, lint Lint) (*Result, error) {
	tree, file, err := Parse(path)
	if err != nil {
		return nil, err
	}
	return RunLint(tree, file, lint), nil
}

func RunLint(tree *ast.File, file *File, lint Lint) *Result {
	if !lint.CheckApplies(tree, file) {
		return nil
	}
	return lint.Lint(tree, file)
}

func RunLints(path string, lints []Lint) ([]*Result, error) {
	tree, file, err := Parse(path)
	if err != nil {
		return nil, err
	}
	var results []*Result
	for _, lint := range lints {
		if result := RunLint(tree, file, lint); result != nil {
			results = append(results, result)
		}
	}
	return results, nil
}
