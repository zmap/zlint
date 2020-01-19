package lint

var (
	// Lints is a map of all known lints by name. Add a Lint to the map by calling
	// RegisterLint.
	Lints = make(map[string]*Lint)
)

// RegisterLint must be called once for each lint to be excuted. Duplicate lint
// names are squashed. Normally, RegisterLint is called during init().
func RegisterLint(l *Lint) {
	if err := l.Lint.Initialize(); err != nil {
		panic("could not initialize lint: " + l.Name + ": " + err.Error())
	}
	Lints[l.Name] = l
}
