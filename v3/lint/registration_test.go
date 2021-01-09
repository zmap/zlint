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
	"errors"
	"reflect"
	"regexp"
	"sort"
	"testing"

	"github.com/zmap/zcrypto/x509"
)

func TestAllLintsHaveNameDescriptionSource(t *testing.T) {
	for _, name := range GlobalRegistry().Names() {
		lint := GlobalRegistry().ByName(name)
		if lint.Name == "" {
			t.Errorf("lint %s has empty name", name)
		}
		if lint.Description == "" {
			t.Errorf("lint %s has empty description", name)
		}
		if lint.Citation == "" {
			t.Errorf("lint %s has empty citation", name)
		}
	}
}

func TestAllLintsHaveSource(t *testing.T) {
	for _, name := range globalRegistry.Names() {
		lint := GlobalRegistry().ByName(name)
		if lint.Source == UnknownLintSource {
			t.Errorf("lint %s has unknown source", name)
		}
	}
}

func TestFilterOptionsEmpty(t *testing.T) {
	opts := FilterOptions{}
	if !opts.Empty() {
		t.Errorf("Empty FilterOptions wasn't Empty()")
	}
	opts.IncludeNames = []string{"whatever"}
	if opts.Empty() {
		t.Errorf("Non-empty FilterOptions was Empty()")
	}
}

type mockLint struct {
	initErr error
}

func (m mockLint) Initialize() error {
	return m.initErr
}

func (m mockLint) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (m mockLint) Execute(c *x509.Certificate) *LintResult {
	return nil
}

func TestRegister(t *testing.T) {
	egLint := &Lint{
		Name:   "mockLint",
		Lint:   &mockLint{},
		Source: Community,
	}
	dupeReg := NewRegistry()
	_ = dupeReg.register(egLint, true)

	badInitErr := errors.New("mock init error")
	badInitLint := &Lint{
		Name:   "badInitLint",
		Lint:   &mockLint{badInitErr},
		Source: Community,
	}

	testCases := []struct {
		name          string
		lint          *Lint
		init          bool
		registry      *registryImpl
		expectErr     error
		expectNames   []string
		expectSources SourceList
	}{
		{
			name:      "nil lint",
			lint:      nil,
			expectErr: errNilLint,
		},
		{
			name:      "nil lint ptr",
			lint:      &Lint{},
			expectErr: errNilLintPtr,
		},
		{
			name: "empty name",
			lint: &Lint{
				Lint: &mockLint{},
			},
			expectErr: errEmptyName,
		},
		{
			name:      "duplicate name",
			lint:      egLint,
			registry:  dupeReg,
			expectErr: &errDuplicateName{egLint.Name},
		},
		{
			name:      "bad init with initialize",
			lint:      badInitLint,
			init:      true,
			expectErr: &errBadInit{badInitLint.Name, badInitErr},
		},
		{
			name:          "bad init with no initialize",
			lint:          badInitLint,
			init:          false,
			expectNames:   []string{badInitLint.Name},
			expectSources: SourceList{badInitLint.Source},
		},
		{
			name: "good lint register",
			lint: &Lint{
				Name:   "goodLint",
				Lint:   &mockLint{},
				Source: MozillaRootStorePolicy,
			},
			registry:      dupeReg,
			expectNames:   []string{"goodLint", egLint.Name},
			expectSources: SourceList{egLint.Source, MozillaRootStorePolicy},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var reg *registryImpl
			if tc.registry == nil {
				reg = NewRegistry()
			} else {
				reg = tc.registry
			}

			err := reg.register(tc.lint, tc.init)
			if err == nil && tc.expectErr != nil {
				t.Errorf("expected err %v, got nil", tc.expectErr)
			} else if err != nil && err.Error() != tc.expectErr.Error() {
				t.Errorf("expected err %v got %v", tc.expectErr, err)
			} else if err == nil {
				if !reflect.DeepEqual(reg.Names(), tc.expectNames) {
					t.Errorf("expected names %v, got %v", tc.expectNames, reg.Names())
				}
				sources := reg.Sources()
				sort.Sort(sources)
				if !reflect.DeepEqual(sources, tc.expectSources) {
					t.Errorf("expected sources %v, got %v", tc.expectSources, sources)
				}
			}
		})
	}
}

func TestRegistryFilter(t *testing.T) {
	testLint := func(name string, source LintSource) *Lint {
		return &Lint{
			Name:   name,
			Source: source,
			Lint:   &mockLint{},
		}
	}
	mustRegister := func(r *registryImpl, l *Lint) {
		if err := r.register(l, true); err != nil {
			t.Fatalf("failed to register %v", err)
		}
	}

	// Create a registry and add some test lints
	registry := NewRegistry()

	mustRegister(registry, testLint("e_mp_example1", MozillaRootStorePolicy))
	mustRegister(registry, testLint("w_mp_example2", MozillaRootStorePolicy))
	mustRegister(registry, testLint("n_mp_example3", MozillaRootStorePolicy))
	mustRegister(registry, testLint("e_z_example1", Community))
	mustRegister(registry, testLint("e_rfc_example1", RFC5280))
	mustRegister(registry, testLint("w_rfc_example2", RFC5280))

	onlyWarnRegex := regexp.MustCompile(`^w\_.*`)

	// Up front, test that invalid FilterOptions provokes an err
	_, err := registry.Filter(FilterOptions{
		NameFilter:   onlyWarnRegex,
		IncludeNames: []string{"e_mp_example_1"},
	})
	if err == nil {
		t.Errorf("expected err from invalid FilterOptions, got nil")
	}

	testCases := []struct {
		name              string
		opts              FilterOptions
		expectedLintNames []string
		expectedSources   SourceList
	}{
		{
			name: "Empty filter options",
			expectedLintNames: []string{
				"e_mp_example1", "e_rfc_example1", "e_z_example1", "n_mp_example3", "w_mp_example2", "w_rfc_example2",
			},
			expectedSources: SourceList{
				Community, MozillaRootStorePolicy, RFC5280,
			},
		},
		{
			name: "Filter by NameFilter only",
			opts: FilterOptions{
				NameFilter: onlyWarnRegex,
			},
			expectedLintNames: []string{
				"w_mp_example2", "w_rfc_example2",
			},
			expectedSources: SourceList{
				MozillaRootStorePolicy, RFC5280,
			},
		},
		{
			name: "Filter by IncludeNames only",
			opts: FilterOptions{
				IncludeNames: []string{
					"e_rfc_example1", "w_mp_example2",
				},
			},
			expectedLintNames: []string{
				"e_rfc_example1", "w_mp_example2",
			},
			expectedSources: SourceList{
				MozillaRootStorePolicy, RFC5280,
			},
		},
		{
			name: "Filter by ExcludeNames only",
			opts: FilterOptions{
				ExcludeNames: []string{
					"e_rfc_example1", "w_mp_example2",
				},
			},
			expectedLintNames: []string{
				"e_mp_example1", "e_z_example1", "n_mp_example3", "w_rfc_example2",
			},
			expectedSources: SourceList{
				Community, MozillaRootStorePolicy, RFC5280,
			},
		},
		{
			name: "Filter by ExcludeNames and IncludeNames",
			opts: FilterOptions{
				ExcludeNames: []string{
					"e_rfc_example1", "w_mp_example2",
				},
				IncludeNames: []string{
					"e_rfc_example1", "e_z_example1",
				},
			},
			expectedLintNames: []string{
				"e_z_example1",
			},
			expectedSources: SourceList{
				Community,
			},
		},
		{
			name: "Filter by IncludeSources only",
			opts: FilterOptions{
				IncludeSources: SourceList{
					Community, RFC5280,
				},
			},
			expectedLintNames: []string{
				"e_rfc_example1", "e_z_example1", "w_rfc_example2",
			},
			expectedSources: SourceList{
				Community, RFC5280,
			},
		},
		{
			name: "Filter by ExcludeSources only",
			opts: FilterOptions{
				ExcludeSources: SourceList{
					RFC5280,
				},
			},
			expectedLintNames: []string{
				"e_mp_example1", "e_z_example1", "n_mp_example3", "w_mp_example2",
			},
			expectedSources: SourceList{
				Community, MozillaRootStorePolicy,
			},
		},
		{
			name: "Filter by IncludeSources and ExcludeSources",
			opts: FilterOptions{
				ExcludeSources: SourceList{
					RFC5280,
				},
				IncludeSources: SourceList{
					Community,
				},
			},
			expectedLintNames: []string{
				"e_z_example1",
			},
			expectedSources: SourceList{
				Community,
			},
		},
		{
			name: "Filter by IncludeSources, ExcludeSources and NameFilter",
			opts: FilterOptions{
				NameFilter: onlyWarnRegex,
				ExcludeSources: SourceList{
					Community,
				},
				IncludeSources: SourceList{
					MozillaRootStorePolicy,
					RFC5280,
				},
			},
			expectedLintNames: []string{
				"w_mp_example2", "w_rfc_example2",
			},
			expectedSources: SourceList{
				MozillaRootStorePolicy, RFC5280,
			},
		},
		{
			name: "Filter by IncludeSources, ExcludeSources, IncludeNames and ExcludeNames",
			opts: FilterOptions{
				ExcludeSources: SourceList{
					Community,
				},
				IncludeSources: SourceList{
					MozillaRootStorePolicy,
					RFC5280,
				},
				ExcludeNames: []string{"e_mp_example1"},
				IncludeNames: []string{"e_rfc_example1", "w_mp_example2"},
			},
			expectedLintNames: []string{
				"e_rfc_example1", "w_mp_example2",
			},
			expectedSources: SourceList{
				MozillaRootStorePolicy, RFC5280,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := registry.Filter(tc.opts)
			if err != nil {
				t.Fatalf("Filter returned err for %v", tc.opts)
			}

			if !reflect.DeepEqual(result.Names(), tc.expectedLintNames) {
				t.Errorf("expected post-Filter Names %v got %v", tc.expectedLintNames, result.Names())
			}

			sources := result.Sources()
			sort.Sort(sources)
			if !reflect.DeepEqual(sources, tc.expectedSources) {
				t.Errorf("expected post-Filter Sources %v got %v", tc.expectedSources, sources)
			}
		})
	}
}
