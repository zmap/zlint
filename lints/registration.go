package lint

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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
