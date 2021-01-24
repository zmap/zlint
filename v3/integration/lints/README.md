# Linting the Linter

This directory contains a collection of Golang code linters that are intended to be very specific to ZLint itself.

# Running

```bash
go run main.go <path to code directory>
```

The linter will walk the given directory recursively and attempt to parse and lint each Go file it comes accross.

In order to extend this custom linter, write a new Go file in the `lints` directory which contains a struct that implements the following interface.

# Extending

```go
type Lint interface {
    Lint(tree *ast.File, file *File) *Result
    CheckApplies(tree *ast.File, file *File) bool
}
```

Then go in to `main.go` and add a pointer to your lint to the `Linters` slice.

```go
var Linters = []lint.Lint{
    &lints.InitFirst{},
    &lints.MySuperCoolLint{}
}
```
