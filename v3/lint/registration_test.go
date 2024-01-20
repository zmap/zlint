package lint

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
	"reflect"
	"regexp"
	"sort"
	"testing"

	"github.com/zmap/zcrypto/x509"
)

func TestAllLintsHaveValidMeta(t *testing.T) {
	checkMeta := func(meta LintMetadata) {
		if meta.Name == "" {
			t.Errorf("lint %s has empty name", meta.Name)
		}
		if meta.Description == "" {
			t.Errorf("lint %s has empty description", meta.Name)
		}
		if meta.Citation == "" {
			t.Errorf("lint %s has empty citation", meta.Name)
		}
		if meta.Source == UnknownLintSource {
			t.Errorf("lint %s has unknown source", meta.Name)
		}
	}
	for _, lint := range globalRegistry.certificateLints.lints {
		checkMeta(lint.LintMetadata)
	}
	for _, lint := range globalRegistry.revocationListLints.lints {
		checkMeta(lint.LintMetadata)
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

type mockLint struct{}

func (m mockLint) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (m mockLint) Execute(c *x509.Certificate) *LintResult {
	return nil
}

type mockRevocationListLint struct{}

func (m mockRevocationListLint) CheckApplies(c *x509.RevocationList) bool {
	return true
}

func (m mockRevocationListLint) Execute(c *x509.RevocationList) *LintResult {
	return nil
}

func TestRegister(t *testing.T) {
	egLint := &Lint{
		Name:   "mockLint",
		Lint:   func() LintInterface { return &mockLint{} },
		Source: Community,
	}
	dupeReg := NewRegistry()
	_ = dupeReg.register(egLint)

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
			name: "nil lint ptr",
			lint: &Lint{
				Lint: func() LintInterface { return nil },
			},
			expectErr: errNilLintPtr,
		},
		{
			name: "empty name",
			lint: &Lint{
				Lint: func() LintInterface { return &mockLint{} },
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
			name: "good lint register",
			lint: &Lint{
				Name:   "goodLint",
				Lint:   func() LintInterface { return &mockLint{} },
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

			err := reg.register(tc.lint)
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

func TestRegistryLookupEngine(t *testing.T) {
	expectedNames := []string{
		"A-mockCertificateLint",
		"B-mockLint",
		"C-mockRevocationListLint",
	}

	expectedSources := []LintSource{
		Community,
		RFC3279,
		RFC8813,
	}

	egCertificateLint := &CertificateLint{
		LintMetadata: LintMetadata{
			Name:   "A-mockCertificateLint",
			Source: Community,
		},
		Lint: func() CertificateLintInterface { return &mockLint{} },
	}

	egLint := &Lint{
		Name:   "B-mockLint",
		Lint:   func() LintInterface { return &mockLint{} },
		Source: RFC8813, // arbitrary value for testing
	}

	egRevocationListLint := &RevocationListLint{
		LintMetadata: LintMetadata{
			Name:   "C-mockRevocationListLint",
			Source: RFC3279, // arbitrary value for testing
		},
		Lint: func() RevocationListLintInterface { return &mockRevocationListLint{} },
	}

	registry := NewRegistry()
	if err := registry.register(egLint); err != nil {
		t.Fatalf("registry.register failed: %v", err)
	}
	if err := registry.registerCertificateLint(egCertificateLint); err != nil {
		t.Fatalf("registry.registerCertificateLint failed: %v", err)
	}
	if err := registry.registerRevocationListLint(egRevocationListLint); err != nil {
		t.Fatalf("registry.registerRevocationListLint failed: %v", err)
	}
	t.Run("lint names are correct and sorted", func(t *testing.T) {
		if !reflect.DeepEqual(registry.Names(), expectedNames) {
			t.Fatalf("expected lint names: %v, got: %v", registry.Names(), expectedNames)
		}
	})

	t.Run("sources are valid", func(t *testing.T) {
		sources := registry.Sources()
		sort.Sort(sources)
		for i, source := range sources {
			if source != expectedSources[i] {
				t.Fatalf("expected source names: %v, got: %v", sources, expectedSources)
			}
		}
	})

	t.Run("stores contain correct lints", func(t *testing.T) {
		testCases := []struct {
			name                string
			deprecatedStore     bool
			certificateStore    bool
			revocationListStore bool
		}{
			{
				name:                "A-mockCertificateLint",
				deprecatedStore:     true,
				certificateStore:    true,
				revocationListStore: false,
			},
			{
				name:                "B-mockLint",
				deprecatedStore:     true,
				certificateStore:    true,
				revocationListStore: false,
			},
			{
				name:                "C-mockRevocationListLint",
				deprecatedStore:     false,
				certificateStore:    false,
				revocationListStore: true,
			},
		}

		for _, tc := range testCases {
			{
				lint := registry.ByName(tc.name)
				if (lint != nil) != tc.deprecatedStore {
					t.Fatalf("expected lint %s to be %t (true = present, false = absent) in deprecated store", tc.name, tc.deprecatedStore)
				}
			}
			{
				lint := registry.CertificateLints().ByName(tc.name)
				if (lint != nil) != tc.certificateStore {
					t.Fatalf("expected lint %s to be %t (true = present, false = absent) in certificate store", tc.name, tc.certificateStore)
				}
			}
			{
				lint := registry.RevocationListLints().ByName(tc.name)
				if (lint != nil) != tc.revocationListStore {
					t.Fatalf("expected lint %s to be %t (true = present, false = absent) in revocationList store", tc.name, tc.revocationListStore)
				}
			}
		}
	})
}

func TestRegistryFilter(t *testing.T) {
	testLint := func(name string, source LintSource) *Lint {
		return &Lint{
			Name:   name,
			Source: source,
			Lint:   func() LintInterface { return &mockLint{} },
		}
	}
	mustRegister := func(r *registryImpl, l *Lint) {
		if err := r.register(l); err != nil {
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
